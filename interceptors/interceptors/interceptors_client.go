// Code generated by goa v3.19.1, DO NOT EDIT.
//
// interceptors example client interceptors
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

// InterceptorsClientInterceptors implements the client interceptors for the interceptors service.
type InterceptorsClientInterceptors struct {
}

// NewInterceptorsClientInterceptors creates a new client interceptor for the interceptors service.
func NewInterceptorsClientInterceptors() *InterceptorsClientInterceptors {
	return &InterceptorsClientInterceptors{}
}

// Client-side interceptor which writes the tenant ID to the signed JWT
// contained in the Authorization header
func (i *InterceptorsClientInterceptors) EncodeTenant(ctx context.Context, info *interceptors.EncodeTenantInfo, next goa.Endpoint) (any, error) {
	log.Printf(ctx, "[EncodeTenant] Sending request: %v", info.RawPayload())
	resp, err := next(ctx, info.RawPayload())
	if err != nil {
		log.Printf(ctx, "[EncodeTenant] Error: %v", err)
		return nil, err
	}
	log.Printf(ctx, "[EncodeTenant] Received response: %v", resp)
	return resp, nil
}

// Client-side interceptor which implements smart retry logic with exponential
// backoff
func (i *InterceptorsClientInterceptors) Retry(ctx context.Context, info *interceptors.RetryInfo, next goa.Endpoint) (any, error) {
	log.Printf(ctx, "[Retry] Sending request: %v", info.RawPayload())
	resp, err := next(ctx, info.RawPayload())
	if err != nil {
		log.Printf(ctx, "[Retry] Error: %v", err)
		return nil, err
	}
	log.Printf(ctx, "[Retry] Received response: %v", resp)
	return resp, nil
}

// Server-side and client-side interceptor that adds trace context to the
// bidirectional stream payload
func (i *InterceptorsClientInterceptors) TraceBidirectionalStream(ctx context.Context, info *interceptors.TraceBidirectionalStreamInfo, next goa.Endpoint) (any, error) {
	log.Printf(ctx, "[TraceBidirectionalStream] Sending request: %v", info.RawPayload())
	resp, err := next(ctx, info.RawPayload())
	if err != nil {
		log.Printf(ctx, "[TraceBidirectionalStream] Error: %v", err)
		return nil, err
	}
	log.Printf(ctx, "[TraceBidirectionalStream] Received response: %v", resp)
	return resp, nil
}
