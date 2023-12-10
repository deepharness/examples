// Code generated by goa v3.14.1, DO NOT EDIT.
//
// default_service service
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design

package defaultservice

import (
	"context"

	"goa.design/goa/v3/security"
)

// Service is the default_service service interface.
type Service interface {
	// The default service default_method is secured with basic authentication
	Default(context.Context, *DefaultPayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// BasicAuth implements the authorization logic for the Basic security scheme.
	BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "default_service"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"default"}

// DefaultPayload is the payload type of the default_service service default
// method.
type DefaultPayload struct {
	Username string
	Password string
}
