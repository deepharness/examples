// Code generated by goa v3.14.1, DO NOT EDIT.
//
// session endpoints
//
// Command:
// $ goa gen goa.design/examples/cookies/design

package session

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "session" service endpoints.
type Endpoints struct {
	CreateSession goa.Endpoint
	UseSession    goa.Endpoint
}

// NewEndpoints wraps the methods of the "session" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		CreateSession: NewCreateSessionEndpoint(s),
		UseSession:    NewUseSessionEndpoint(s),
	}
}

// Use applies the given middleware to all the "session" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.CreateSession = m(e.CreateSession)
	e.UseSession = m(e.UseSession)
}

// NewCreateSessionEndpoint returns an endpoint function that calls the method
// "create_session" of service "session".
func NewCreateSessionEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreateSessionPayload)
		return s.CreateSession(ctx, p)
	}
}

// NewUseSessionEndpoint returns an endpoint function that calls the method
// "use_session" of service "session".
func NewUseSessionEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UseSessionPayload)
		return s.UseSession(ctx, p)
	}
}
