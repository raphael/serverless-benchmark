// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// recorder HTTP client types
//
// Command:
// $ goa gen github.com/raphael/recorder/design

package client

import (
	recordersvc "github.com/raphael/recorder/gen/recorder"
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
}

// NewRecordDataRequestBody builds the HTTP request body from the payload of
// the "record-data" endpoint of the "recorder" service.
func NewRecordDataRequestBody(p *recordersvc.Datapoint) *RecordDataRequestBody {
	body := &RecordDataRequestBody{
		Service: p.Service,
		Value:   p.Value,
		Name:    p.Name,
	}
	return body
}
