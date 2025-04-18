// Code generated by goa v3.20.1, DO NOT EDIT.
//
// probes endpoints
//
// Command:
// $ goa gen github.com/meier-christoph/todo-backend-golang-goa/design

package probes

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "probes" service endpoints.
type Endpoints struct {
	Healthy goa.Endpoint
}

// NewEndpoints wraps the methods of the "probes" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Healthy: NewHealthyEndpoint(s),
	}
}

// Use applies the given middleware to all the "probes" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Healthy = m(e.Healthy)
}

// NewHealthyEndpoint returns an endpoint function that calls the method
// "healthy" of service "probes".
func NewHealthyEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return nil, s.Healthy(ctx)
	}
}
