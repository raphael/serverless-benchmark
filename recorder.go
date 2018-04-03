package recorder

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"cloud.google.com/go/storage"
	recorder "github.com/raphael/recorder/gen/recorder"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"
)

type (
	// recorder service implementation
	recorderSvc struct{}
)

// NewRecorder returns the recorder service implementation.
func NewRecorder() recorder.Service {
	return &recorderSvc{}
}

// RecordData creates a new datapoint.
func (s *recorderSvc) RecordData(ctx context.Context, p *recorder.Datapoint) error {
	all, err := read(ctx, p.Service, p.Name)
	if err != nil {
		return err
	}
	all = append(all, p)
	return write(ctx, p.Service, p.Name, all)
}

// List lists all datapoints for the given service and name.
func (s *recorderSvc) List(ctx context.Context, p *recorder.Series) ([]*recorder.Datapoint, error) {
	return read(ctx, p.Service, p.Name)
}

func read(ctx context.Context, service, name string) ([]*recorder.Datapoint, error) {
	var pts []*recorder.Datapoint
	err := withBucket(ctx, func(bucket *storage.BucketHandle) error {
		fn := filename(service, name)
		rc, err := bucket.Object(fn).NewReader(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "object doesn't exist") {
				return nil
			}
			return fmt.Errorf("readall: unable to open file %q: %v", fn, err)
		}
		defer rc.Close()
		slurp, err := ioutil.ReadAll(rc)
		if err != nil {
			return fmt.Errorf("readall: unable to read data from file %q: %v", fn, err)
		}
		if len(slurp) > 0 {
			if err := json.Unmarshal(slurp, &pts); err != nil {
				return fmt.Errorf("readall: failed to deserialize content of file %q: %v", fn, err)
			}
		}
		return nil
	})
	if err != nil {
		log.Errorf(ctx, err.Error())
	}
	return pts, err
}

func write(ctx context.Context, service, name string, pts []*recorder.Datapoint) error {
	err := withBucket(ctx, func(bucket *storage.BucketHandle) error {
		fn := filename(service, name)
		wc := bucket.Object(fn).NewWriter(ctx)
		defer wc.Close()
		wc.ContentType = "text/plain"
		js, err := json.Marshal(pts)
		if err != nil {
			return err
		}
		_, err = wc.Write(js)
		return err
	})
	if err != nil {
		log.Errorf(ctx, err.Error())
	}
	return err
}

func withBucket(ctx context.Context, fn func(*storage.BucketHandle) error) error {
	bucket, err := file.DefaultBucketName(ctx)
	if err != nil {
		return fmt.Errorf("readall: failed to get default GCS bucket name: %v", err)
	}
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()
	return fn(client.Bucket(bucket))
}

func filename(service, name string) string {
	return fmt.Sprintf("%s-%s.txt", service, name)
}
