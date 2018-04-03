// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// recorder HTTP server
//
// Command:
// $ goa gen github.com/raphael/recorder/design

package server

import (
	"context"
	"net/http"

	recordersvc "github.com/raphael/recorder/gen/recorder"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Server lists the recorder service endpoint HTTP handlers.
type Server struct {
	Mounts     []*MountPoint
	RecordData http.Handler
	List       http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the recorder service endpoints.
func New(
	e *recordersvc.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"RecordData", "POST", "/data"},
			{"List", "GET", "/data"},
		},
		RecordData: NewRecordDataHandler(e.RecordData, mux, dec, enc, eh),
		List:       NewListHandler(e.List, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "recorder" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.RecordData = m(s.RecordData)
	s.List = m(s.List)
}

// Mount configures the mux to serve the recorder endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountRecordDataHandler(mux, h.RecordData)
	MountListHandler(mux, h.List)
}

// MountRecordDataHandler configures the mux to serve the "recorder" service
// "record-data" endpoint.
func MountRecordDataHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/data", f)
}

// NewRecordDataHandler creates a HTTP handler which loads the HTTP request and
// calls the "recorder" service "record-data" endpoint.
func NewRecordDataHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRecordDataRequest(mux, dec)
		encodeResponse = EncodeRecordDataResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "record-data")
		ctx = context.WithValue(ctx, goa.ServiceKey, "recorder")
		payload, err := decodeRequest(r)
		if err != nil {
			eh(ctx, w, err)
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
				return
			}
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountListHandler configures the mux to serve the "recorder" service "list"
// endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/data", f)
}

// NewListHandler creates a HTTP handler which loads the HTTP request and calls
// the "recorder" service "list" endpoint.
func NewListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeListRequest(mux, dec)
		encodeResponse = EncodeListResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list")
		ctx = context.WithValue(ctx, goa.ServiceKey, "recorder")
		payload, err := decodeRequest(r)
		if err != nil {
			eh(ctx, w, err)
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
				return
			}
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
