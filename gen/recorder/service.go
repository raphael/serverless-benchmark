// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// recorder service
//
// Command:
// $ goa gen github.com/raphael/recorder/design

package recordersvc

import (
	"context"
)

// Service is the recorder service interface.
type Service interface {
	// RecordData creates a new datapoint.
	RecordData(context.Context, *Datapoint) error
	// List lists all recorded datapoints.
	List(context.Context, *Series) ([]float64, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "recorder"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"record-data", "list"}

// Datapoint is the payload type of the recorder service record-data method.
type Datapoint struct {
	// Service that created datapoint.
	Service string
	// Datapoint value.
	Value float64
	// Name is the name of the datapoint.
	Name string
}

// Series is the payload type of the recorder service list method.
type Series struct {
	// Service that created datapoint.
	Service string
	// Name is the name of the datapoint.
	Name string
}
