// Code generated by goa v3.7.3, DO NOT EDIT.
//
// chatter endpoints
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package chatter

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "chatter" service endpoints.
type Endpoints struct {
	Login     goa.Endpoint
	Echoer    goa.Endpoint
	Listener  goa.Endpoint
	Summary   goa.Endpoint
	Subscribe goa.Endpoint
	History   goa.Endpoint
}

// EchoerEndpointInput holds both the payload and the server stream of the
// "echoer" method.
type EchoerEndpointInput struct {
	// Payload is the method payload.
	Payload *EchoerPayload
	// Stream is the server stream used by the "echoer" method to send data.
	Stream EchoerServerStream
}

// ListenerEndpointInput holds both the payload and the server stream of the
// "listener" method.
type ListenerEndpointInput struct {
	// Payload is the method payload.
	Payload *ListenerPayload
	// Stream is the server stream used by the "listener" method to send data.
	Stream ListenerServerStream
}

// SummaryEndpointInput holds both the payload and the server stream of the
// "summary" method.
type SummaryEndpointInput struct {
	// Payload is the method payload.
	Payload *SummaryPayload
	// Stream is the server stream used by the "summary" method to send data.
	Stream SummaryServerStream
}

// SubscribeEndpointInput holds both the payload and the server stream of the
// "subscribe" method.
type SubscribeEndpointInput struct {
	// Payload is the method payload.
	Payload *SubscribePayload
	// Stream is the server stream used by the "subscribe" method to send data.
	Stream SubscribeServerStream
}

// HistoryEndpointInput holds both the payload and the server stream of the
// "history" method.
type HistoryEndpointInput struct {
	// Payload is the method payload.
	Payload *HistoryPayload
	// Stream is the server stream used by the "history" method to send data.
	Stream HistoryServerStream
}

// NewEndpoints wraps the methods of the "chatter" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Login:     NewLoginEndpoint(s, a.BasicAuth),
		Echoer:    NewEchoerEndpoint(s, a.JWTAuth),
		Listener:  NewListenerEndpoint(s, a.JWTAuth),
		Summary:   NewSummaryEndpoint(s, a.JWTAuth),
		Subscribe: NewSubscribeEndpoint(s, a.JWTAuth),
		History:   NewHistoryEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "chatter" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Login = m(e.Login)
	e.Echoer = m(e.Echoer)
	e.Listener = m(e.Listener)
	e.Summary = m(e.Summary)
	e.Subscribe = m(e.Subscribe)
	e.History = m(e.History)
}

// NewLoginEndpoint returns an endpoint function that calls the method "login"
// of service "chatter".
func NewLoginEndpoint(s Service, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LoginPayload)
		var err error
		sc := security.BasicScheme{
			Name:           "basic",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authBasicFn(ctx, p.User, p.Password, &sc)
		if err != nil {
			return nil, err
		}
		return s.Login(ctx, p)
	}
}

// NewEchoerEndpoint returns an endpoint function that calls the method
// "echoer" of service "chatter".
func NewEchoerEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*EchoerEndpointInput)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"stream:read", "stream:write"},
			RequiredScopes: []string{"stream:write"},
		}
		ctx, err = authJWTFn(ctx, ep.Payload.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Echoer(ctx, ep.Payload, ep.Stream)
	}
}

// NewListenerEndpoint returns an endpoint function that calls the method
// "listener" of service "chatter".
func NewListenerEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*ListenerEndpointInput)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"stream:read", "stream:write"},
			RequiredScopes: []string{"stream:write"},
		}
		ctx, err = authJWTFn(ctx, ep.Payload.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Listener(ctx, ep.Payload, ep.Stream)
	}
}

// NewSummaryEndpoint returns an endpoint function that calls the method
// "summary" of service "chatter".
func NewSummaryEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*SummaryEndpointInput)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"stream:read", "stream:write"},
			RequiredScopes: []string{"stream:write"},
		}
		ctx, err = authJWTFn(ctx, ep.Payload.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Summary(ctx, ep.Payload, ep.Stream)
	}
}

// NewSubscribeEndpoint returns an endpoint function that calls the method
// "subscribe" of service "chatter".
func NewSubscribeEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*SubscribeEndpointInput)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"stream:read", "stream:write"},
			RequiredScopes: []string{"stream:write"},
		}
		ctx, err = authJWTFn(ctx, ep.Payload.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.Subscribe(ctx, ep.Payload, ep.Stream)
	}
}

// NewHistoryEndpoint returns an endpoint function that calls the method
// "history" of service "chatter".
func NewHistoryEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*HistoryEndpointInput)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"stream:read", "stream:write"},
			RequiredScopes: []string{"stream:read"},
		}
		ctx, err = authJWTFn(ctx, ep.Payload.Token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.History(ctx, ep.Payload, ep.Stream)
	}
}
