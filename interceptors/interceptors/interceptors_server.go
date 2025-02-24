// Code generated by goa v3.19.1, DO NOT EDIT.
//
// interceptors example server interceptors
//
// Command:
// $ goa example goa.design/examples/interceptors/design

package interceptors

import (
	"context"

	"goa.design/clue/log"
	interceptors "goa.design/examples/interceptors/gen/interceptors"
	goa "goa.design/goa/v3/pkg"
)

// InterceptorsServerInterceptors implements the server interceptor for the interceptors service.
type InterceptorsServerInterceptors struct {
}

// NewInterceptorsServerInterceptors creates a new server interceptor for the interceptors service.
func NewInterceptorsServerInterceptors() *InterceptorsServerInterceptors {
	return &InterceptorsServerInterceptors{}
}

// Server-side interceptor which implements a transparent cache for loaded
// records
func (i *InterceptorsServerInterceptors) Cache(ctx context.Context, info *interceptors.CacheInfo, next goa.Endpoint) (any, error) {
	log.Printf(ctx, "[Cache] Processing request: %v", info.RawPayload())
	resp, err := next(ctx, info.RawPayload())
	if err != nil {
		log.Printf(ctx, "[Cache] Error: %v", err)
		return nil, err
	}
	log.Printf(ctx, "[Cache] Response: %v", resp)
	return resp, nil
}

// Server-side interceptor that validates JWT token and tenant ID
func (i *InterceptorsServerInterceptors) JWTAuth(ctx context.Context, info *interceptors.JWTAuthInfo, next goa.Endpoint) (any, error) {
	log.Printf(ctx, "[JWTAuth] Processing request: %v", info.RawPayload())
	resp, err := next(ctx, info.RawPayload())
	if err != nil {
		log.Printf(ctx, "[JWTAuth] Error: %v", err)
		return nil, err
	}
	log.Printf(ctx, "[JWTAuth] Response: %v", resp)
	return resp, nil
}

// Server-side interceptor that provides comprehensive request/response audit
// logging
func (i *InterceptorsServerInterceptors) RequestAudit(ctx context.Context, info *interceptors.RequestAuditInfo, next goa.Endpoint) (any, error) {
	log.Printf(ctx, "[RequestAudit] Processing request: %v", info.RawPayload())
	resp, err := next(ctx, info.RawPayload())
	if err != nil {
		log.Printf(ctx, "[RequestAudit] Error: %v", err)
		return nil, err
	}
	log.Printf(ctx, "[RequestAudit] Response: %v", resp)
	return resp, nil
}

// Server-side interceptor which sets the context deadline for the request
func (i *InterceptorsServerInterceptors) SetDeadline(ctx context.Context, info *interceptors.SetDeadlineInfo, next goa.Endpoint) (any, error) {
	log.Printf(ctx, "[SetDeadline] Processing request: %v", info.RawPayload())
	resp, err := next(ctx, info.RawPayload())
	if err != nil {
		log.Printf(ctx, "[SetDeadline] Error: %v", err)
		return nil, err
	}
	log.Printf(ctx, "[SetDeadline] Response: %v", resp)
	return resp, nil
}

// Server-side and client-side interceptor that adds trace context to the
// bidirectional stream payload
func (i *InterceptorsServerInterceptors) TraceBidirectionalStream(ctx context.Context, info *interceptors.TraceBidirectionalStreamInfo, next goa.Endpoint) (any, error) {
	log.Printf(ctx, "[TraceBidirectionalStream] Processing request: %v", info.RawPayload())
	resp, err := next(ctx, info.RawPayload())
	if err != nil {
		log.Printf(ctx, "[TraceBidirectionalStream] Error: %v", err)
		return nil, err
	}
	log.Printf(ctx, "[TraceBidirectionalStream] Response: %v", resp)
	return resp, nil
}

// Server-side interceptor that adds trace context to the request payload
func (i *InterceptorsServerInterceptors) TraceRequest(ctx context.Context, info *interceptors.TraceRequestInfo, next goa.Endpoint) (any, error) {
	log.Printf(ctx, "[TraceRequest] Processing request: %v", info.RawPayload())
	resp, err := next(ctx, info.RawPayload())
	if err != nil {
		log.Printf(ctx, "[TraceRequest] Error: %v", err)
		return nil, err
	}
	log.Printf(ctx, "[TraceRequest] Response: %v", resp)
	return resp, nil
}
