package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	recorder "github.com/raphael/recorder"
	recordersvcsvr "github.com/raphael/recorder/gen/http/recorder/server"
	recordersvc "github.com/raphael/recorder/gen/recorder"
	goahttp "goa.design/goa/http"
	"goa.design/goa/http/middleware"
	"google.golang.org/appengine"
)

func main() {
	// Define command line flags, add any other flag required to configure
	// the service.
	var (
		addr = flag.String("listen", ":8080", "HTTP listen `address`")
		dbg  = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger and goa log adapter. Replace logger with your own using
	// your log package of choice. The goa.design/middleware/logging/...
	// packages define log adapters for common log packages.
	var (
		adapter middleware.Logger
		logger  *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[recorder] ", log.Ltime)
		adapter = middleware.NewLogger(logger)
	}

	// Create the structs that implement the services.
	var (
		recorderSvc recordersvc.Service
	)
	{
		recorderSvc = recorder.NewRecorder()
	}

	// Wrap the services in endpoints that can be invoked from other
	// services potentially running in different processes.
	var (
		recorderEndpoints *recordersvc.Endpoints
	)
	{
		recorderEndpoints = recordersvc.NewEndpoints(recorderSvc)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		recorderServer *recordersvcsvr.Server
	)
	{
		eh := ErrorHandler(logger)
		recorderServer = recordersvcsvr.New(recorderEndpoints, mux, dec, enc, eh)
	}

	// Configure the mux.
	recordersvcsvr.Mount(mux, recorderServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler
	{
		var h http.Handler = mux
		if *dbg {
			h = middleware.Debug(mux, os.Stdout)(h)
		}
		h = middleware.Log(adapter)(h)
		h = middleware.RequestID()(h)
		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			aectx := appengine.NewContext(r)
			r = r.WithContext(aectx)
			h.ServeHTTP(w, r)
		})
	}

	// Start AppEngine.
	for _, m := range recorderServer.Mounts {
		logger.Printf("method %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	logger.Printf("listening on %s", *addr)
	http.Handle("/", handler)
	appengine.Main()

	logger.Println("exited")
}

// ErrorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func ErrorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
