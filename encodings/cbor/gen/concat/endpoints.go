// Code generated by goa v3.11.2, DO NOT EDIT.
//
// concat endpoints
//
// Command:
// $ goa gen goa.design/examples/encodings/cbor/design -o encodings/cbor

package concat

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "concat" service endpoints.
type Endpoints struct {
	Concat goa.Endpoint
}

// NewEndpoints wraps the methods of the "concat" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Concat: NewConcatEndpoint(s),
	}
}

// Use applies the given middleware to all the "concat" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Concat = m(e.Concat)
}

// NewConcatEndpoint returns an endpoint function that calls the method
// "concat" of service "concat".
func NewConcatEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ConcatPayload)
		return s.Concat(ctx, p)
	}
}
