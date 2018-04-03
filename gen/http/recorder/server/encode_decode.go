// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// recorder HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/raphael/recorder/design

package server

import (
	"context"
	"io"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// EncodeRecordDataResponse returns an encoder for responses returned by the
// recorder record-data endpoint.
func EncodeRecordDataResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeRecordDataRequest returns a decoder for requests sent to the recorder
// record-data endpoint.
func DecodeRecordDataRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body RecordDataRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = body.Validate()
		if err != nil {
			return nil, err
		}

		return NewRecordDataDatapoint(&body), nil
	}
}

// EncodeListResponse returns an encoder for responses returned by the recorder
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]float64)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the recorder list
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			service string
			name    string
			err     error
		)
		service = r.URL.Query().Get("service")
		if service == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("service", "query string"))
		}
		name = r.URL.Query().Get("name")
		if name == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("name", "query string"))
		}
		if err != nil {
			return nil, err
		}

		return NewListSeries(service, name), nil
	}
}
