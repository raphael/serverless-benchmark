// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// recorder client
//
// Command:
// $ goa gen github.com/raphael/recorder/design

package recordersvc

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "recorder" service client.
type Client struct {
	RecordDataEndpoint goa.Endpoint
	ListEndpoint       goa.Endpoint
}

// NewClient initializes a "recorder" service client given the endpoints.
func NewClient(recordData, list goa.Endpoint) *Client {
	return &Client{
		RecordDataEndpoint: recordData,
		ListEndpoint:       list,
	}
}

// RecordData calls the "record-data" endpoint of the "recorder" service.
func (c *Client) RecordData(ctx context.Context, p *Datapoint) (err error) {
	_, err = c.RecordDataEndpoint(ctx, p)
	return
}

// List calls the "list" endpoint of the "recorder" service.
func (c *Client) List(ctx context.Context, p *Series) (res []*Datapoint, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*Datapoint), nil
}
