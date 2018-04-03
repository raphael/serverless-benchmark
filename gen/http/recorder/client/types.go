// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// recorder HTTP client types
//
// Command:
// $ goa gen github.com/raphael/recorder/design

package client

import (
	recordersvc "github.com/raphael/recorder/gen/recorder"
	goa "goa.design/goa"
)

// RecordDataRequestBody is the type of the "recorder" service "record-data"
// endpoint HTTP request body.
type RecordDataRequestBody struct {
	// Service that created datapoint.
	Service string `form:"service" json:"service" xml:"service"`
	// Datapoint value.
	Value float64 `form:"value" json:"value" xml:"value"`
	// Name is the name of the datapoint.
	Name string `form:"name" json:"name" xml:"name"`
	// Labels is an arbitrary set of key/value pairs attached to the event.
	Labels map[string]string `form:"labels,omitempty" json:"labels,omitempty" xml:"labels,omitempty"`
}

// DatapointResponseBody is used to define fields on response body types.
type DatapointResponseBody struct {
	// Service that created datapoint.
	Service *string `form:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
	// Datapoint value.
	Value *float64 `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// Name is the name of the datapoint.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Labels is an arbitrary set of key/value pairs attached to the event.
	Labels map[string]string `form:"labels,omitempty" json:"labels,omitempty" xml:"labels,omitempty"`
}

// NewRecordDataRequestBody builds the HTTP request body from the payload of
// the "record-data" endpoint of the "recorder" service.
func NewRecordDataRequestBody(p *recordersvc.Datapoint) *RecordDataRequestBody {
	body := &RecordDataRequestBody{
		Service: p.Service,
		Value:   p.Value,
		Name:    p.Name,
	}
	if p.Labels != nil {
		body.Labels = make(map[string]string, len(p.Labels))
		for key, val := range p.Labels {
			tk := key
			tv := val
			body.Labels[tk] = tv
		}
	}
	return body
}

// NewListDatapointOK builds a "recorder" service "list" endpoint result from a
// HTTP "OK" response.
func NewListDatapointOK(body []*DatapointResponseBody) []*recordersvc.Datapoint {
	v := make([]*recordersvc.Datapoint, len(body))
	for i, val := range body {
		v[i] = &recordersvc.Datapoint{
			Service: *val.Service,
			Value:   *val.Value,
			Name:    *val.Name,
		}
		if val.Labels != nil {
			v[i].Labels = make(map[string]string, len(val.Labels))
			for key, val := range val.Labels {
				tk := key
				tv := val
				v[i].Labels[tk] = tv
			}
		}
	}
	return v
}

// Validate runs the validations defined on DatapointResponseBody
func (body *DatapointResponseBody) Validate() (err error) {
	if body.Service == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("service", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "body"))
	}
	return
}
