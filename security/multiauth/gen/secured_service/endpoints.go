// Code generated by goa v3.7.7, DO NOT EDIT.
//
// secured_service endpoints
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design -o security/multiauth

package securedservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "secured_service" service endpoints.
type Endpoints struct {
	Signin           goa.Endpoint
	Secure           goa.Endpoint
	DoublySecure     goa.Endpoint
	AlsoDoublySecure goa.Endpoint
}

// NewEndpoints wraps the methods of the "secured_service" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Signin:           NewSigninEndpoint(s, a.BasicAuth),
		Secure:           NewSecureEndpoint(s, a.JWTAuth),
		DoublySecure:     NewDoublySecureEndpoint(s, a.JWTAuth, a.APIKeyAuth),
		AlsoDoublySecure: NewAlsoDoublySecureEndpoint(s, a.JWTAuth, a.APIKeyAuth, a.OAuth2Auth, a.BasicAuth),
	}
}

// Use applies the given middleware to all the "secured_service" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Signin = m(e.Signin)
	e.Secure = m(e.Secure)
	e.DoublySecure = m(e.DoublySecure)
	e.AlsoDoublySecure = m(e.AlsoDoublySecure)
}

// NewSigninEndpoint returns an endpoint function that calls the method
// "signin" of service "secured_service".
func NewSigninEndpoint(s Service, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SigninPayload)
		var err error
		sc := security.BasicScheme{
			Name:           "basic",
			Scopes:         []string{"api:read"},
			RequiredScopes: []string{},
		}
		ctx, err = authBasicFn(ctx, p.Username, p.Password, &sc)
		if err != nil {
			return nil, err
		}
		return s.Signin(ctx, p)
	}
}

// NewSecureEndpoint returns an endpoint function that calls the method
// "secure" of service "secured_service".
func NewSecureEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SecurePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{"api:read"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Secure(ctx, p)
	}
}

// NewDoublySecureEndpoint returns an endpoint function that calls the method
// "doubly_secure" of service "secured_service".
func NewDoublySecureEndpoint(s Service, authJWTFn security.AuthJWTFunc, authAPIKeyFn security.AuthAPIKeyFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DoublySecurePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{"api:read", "api:write"},
		}
		ctx, err = authJWTFn(ctx, p.Token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{"api:read", "api:write"},
			}
			ctx, err = authAPIKeyFn(ctx, p.Key, &sc)
		}
		if err != nil {
			return nil, err
		}
		return s.DoublySecure(ctx, p)
	}
}

// NewAlsoDoublySecureEndpoint returns an endpoint function that calls the
// method "also_doubly_secure" of service "secured_service".
func NewAlsoDoublySecureEndpoint(s Service, authJWTFn security.AuthJWTFunc, authAPIKeyFn security.AuthAPIKeyFunc, authOAuth2Fn security.AuthOAuth2Func, authBasicFn security.AuthBasicFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AlsoDoublySecurePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{"api:read", "api:write"},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err == nil {
			sc := security.APIKeyScheme{
				Name:           "api_key",
				Scopes:         []string{},
				RequiredScopes: []string{"api:read", "api:write"},
			}
			var key string
			if p.Key != nil {
				key = *p.Key
			}
			ctx, err = authAPIKeyFn(ctx, key, &sc)
		}
		if err != nil {
			sc := security.OAuth2Scheme{
				Name:           "oauth2",
				Scopes:         []string{"api:read", "api:write"},
				RequiredScopes: []string{"api:read", "api:write"},
				Flows: []*security.OAuthFlow{
					&security.OAuthFlow{
						Type:             "authorization_code",
						AuthorizationURL: "http://goa.design/authorization",
						TokenURL:         "http://goa.design/token",
						RefreshURL:       "http://goa.design/refresh",
					},
				},
			}
			var token string
			if p.OauthToken != nil {
				token = *p.OauthToken
			}
			ctx, err = authOAuth2Fn(ctx, token, &sc)
			if err == nil {
				sc := security.BasicScheme{
					Name:           "basic",
					Scopes:         []string{"api:read"},
					RequiredScopes: []string{"api:read", "api:write"},
				}
				var user string
				if p.Username != nil {
					user = *p.Username
				}
				var pass string
				if p.Password != nil {
					pass = *p.Password
				}
				ctx, err = authBasicFn(ctx, user, pass, &sc)
			}
		}
		if err != nil {
			return nil, err
		}
		return s.AlsoDoublySecure(ctx, p)
	}
}
