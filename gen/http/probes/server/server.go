// Code generated by goa v3.20.1, DO NOT EDIT.
//
// probes HTTP server
//
// Command:
// $ goa gen github.com/meier-christoph/todo-backend-golang-goa/design

package server

import (
	"context"
	"net/http"

	probes "github.com/meier-christoph/todo-backend-golang-goa/gen/probes"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the probes service endpoint HTTP handlers.
type Server struct {
	Mounts  []*MountPoint
	Healthy http.Handler
	CORS    http.Handler
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

// New instantiates HTTP handlers for all the probes service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *probes.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Healthy", "GET", "/healthy"},
			{"CORS", "OPTIONS", "/healthy"},
		},
		Healthy: NewHealthyHandler(e.Healthy, mux, decoder, encoder, errhandler, formatter),
		CORS:    NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "probes" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Healthy = m(s.Healthy)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return probes.MethodNames[:] }

// Mount configures the mux to serve the probes endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountHealthyHandler(mux, h.Healthy)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the probes endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountHealthyHandler configures the mux to serve the "probes" service
// "healthy" endpoint.
func MountHealthyHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleProbesOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/healthy", f)
}

// NewHealthyHandler creates a HTTP handler which loads the HTTP request and
// calls the "probes" service "healthy" endpoint.
func NewHealthyHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeHealthyResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "healthy")
		ctx = context.WithValue(ctx, goa.ServiceKey, "probes")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service probes.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleProbesOrigin(h)
	mux.Handle("OPTIONS", "/healthy", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 204 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
}

// HandleProbesOrigin applies the CORS response headers corresponding to the
// origin for the service probes.
func HandleProbesOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "*")
				w.Header().Set("Access-Control-Allow-Headers", "*")
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
