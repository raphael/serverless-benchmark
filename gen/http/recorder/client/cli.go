// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// recorder HTTP client CLI support package
//
// Command:
// $ goa gen github.com/raphael/recorder/design

package client

import (
	"encoding/json"
	"fmt"

	recordersvc "github.com/raphael/recorder/gen/recorder"
)

// BuildRecordDataPayload builds the payload for the recorder record-data
// endpoint from CLI flags.
func BuildRecordDataPayload(recorderRecordDataBody string) (*recordersvc.Datapoint, error) {
	var err error
	var body RecordDataRequestBody
	{
		err = json.Unmarshal([]byte(recorderRecordDataBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"labels\": {\n         \"algo\": \"type2\"\n      },\n      \"name\": \"duration\",\n      \"service\": \"lambda\",\n      \"value\": 0.7428001719977185\n   }'")
		}
	}
	if err != nil {
		return nil, err
	}
	v := &recordersvc.Datapoint{
		Service: body.Service,
		Value:   body.Value,
		Name:    body.Name,
	}
	if body.Labels != nil {
		v.Labels = make(map[string]string, len(body.Labels))
		for key, val := range body.Labels {
			tk := key
			tv := val
			v.Labels[tk] = tv
		}
	}
	return v, nil
}

// BuildListPayload builds the payload for the recorder list endpoint from CLI
// flags.
func BuildListPayload(recorderListService string, recorderListName string) (*recordersvc.Series, error) {
	var service string
	{
		service = recorderListService
	}
	var name string
	{
		name = recorderListName
	}
	payload := &recordersvc.Series{
		Service: service,
		Name:    name,
	}
	return payload, nil
}
