package poster

import (
	"context"
	"net/http"

	"github.com/raphael/recorder/gen/http/recorder/client"
	recorder "github.com/raphael/recorder/gen/recorder"
	"goa.design/goa"
	goahttp "goa.design/goa/http"
)

// A Poster posts datapoints to the recorder service.
type Poster goa.Endpoint

// New creates a datapoint poster.
func New(host string) Poster {
	c := client.NewClient("https", host, http.DefaultClient,
		goahttp.RequestEncoder, goahttp.ResponseDecoder, false)
	return Poster(c.RecordData())
}

// Post posts a new datapoint.
func (p Poster) Post(ctx context.Context, point *recorder.Datapoint) error {
	_, err := p(ctx, point)
	return err
}
