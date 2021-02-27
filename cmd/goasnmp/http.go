package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	goasnmp "github.com/rshade/goasnmp/gen/goasnmp"
	goasnmpkitsvr "github.com/rshade/goasnmp/gen/http/goasnmp/kitserver"
	goasnmpsvr "github.com/rshade/goasnmp/gen/http/goasnmp/server"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, goasnmpEndpoints *goasnmp.Endpoints, wg *sync.WaitGroup, errc chan error, logger log.Logger, debug bool) {

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
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
		goasnmpListHandler *kithttp.Server
		goasnmpAddHandler  *kithttp.Server
		goasnmpServer      *goasnmpsvr.Server
	)
	{
		eh := errorHandler(logger)
		goasnmpListHandler = kithttp.NewServer(
			endpoint.Endpoint(goasnmpEndpoints.List),
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			goasnmpkitsvr.EncodeListResponse(enc),
		)
		goasnmpAddHandler = kithttp.NewServer(
			endpoint.Endpoint(goasnmpEndpoints.Add),
			goasnmpkitsvr.DecodeAddRequest(mux, dec),
			goasnmpkitsvr.EncodeAddResponse(enc),
		)
		goasnmpServer = goasnmpsvr.New(goasnmpEndpoints, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	goasnmpkitsvr.MountListHandler(mux, goasnmpListHandler)
	goasnmpkitsvr.MountAddHandler(mux, goasnmpAddHandler)
	goasnmpkitsvr.MountGenHTTPOpenapiJSON(mux)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(logger)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range goasnmpServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Log("info", fmt.Sprintf("HTTP server listening on %q", u.Host))
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Log("info", fmt.Sprintf("shutting down HTTP server at %q", u.Host))

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Log("info", fmt.Sprintf("[%s] ERROR: %s", id, err.Error()))
	}
}
