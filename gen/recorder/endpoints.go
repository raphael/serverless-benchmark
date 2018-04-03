// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// recorder endpoints
//
// Command:
// $ goa gen github.com/raphael/recorder/design

package recordersvc

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "recorder" service endpoints.
type Endpoints struct {
	RecordData goa.Endpoint
	List       goa.Endpoint
}

// NewEndpoints wraps the methods of the "recorder" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		RecordData: NewRecordDataEndpoint(s),
		List:       NewListEndpoint(s),
	}
}

// Use applies the given middleware to all the "recorder" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.RecordData = m(e.RecordData)
	e.List = m(e.List)
}

// NewRecordDataEndpoint returns an endpoint function that calls the method
// "record-data" of service "recorder".
func NewRecordDataEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Datapoint)
		return nil, s.RecordData(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "recorder".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Series)
		return s.List(ctx, p)
	}
}
