// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// recorder HTTP server types
//
// Command:
// $ goa gen github.com/raphael/recorder/design

package server

import (
	recordersvc "github.com/raphael/recorder/gen/recorder"
	goa "goa.design/goa"
)

// RecordDataRequestBody is the type of the "recorder" service "record-data"
// endpoint HTTP request body.
type RecordDataRequestBody struct {
	// Service that created datapoint.
	Service *string `form:"service,omitempty" json:"service,omitempty" xml:"service,omitempty"`
	// Datapoint value.
	Value *float64 `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// Name is the name of the datapoint.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// NewRecordDataDatapoint builds a recorder service record-data endpoint
// payload.
func NewRecordDataDatapoint(body *RecordDataRequestBody) *recordersvc.Datapoint {
	v := &recordersvc.Datapoint{
		Service: *body.Service,
		Value:   *body.Value,
		Name:    *body.Name,
	}
	return v
}

// NewListSeries builds a recorder service list endpoint payload.
func NewListSeries(service string, name string) *recordersvc.Series {
	return &recordersvc.Series{
		Service: service,
		Name:    name,
	}
}

// Validate runs the validations defined on RecordDataRequestBody
func (body *RecordDataRequestBody) Validate() (err error) {
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
