// Code generated by goa v3.14.1, DO NOT EDIT.
//
// secured_service HTTP server types
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design

package server

import (
	securedservice "goa.design/examples/security/multiauth/gen/secured_service"
)

// SigninResponseBody is the type of the "secured_service" service "signin"
// endpoint HTTP response body.
type SigninResponseBody struct {
	// JWT token
	JWT string `form:"jwt" json:"jwt" xml:"jwt"`
	// API Key
	APIKey string `form:"api_key" json:"api_key" xml:"api_key"`
	// OAuth2 token
	OauthToken string `form:"oauth_token" json:"oauth_token" xml:"oauth_token"`
}

// NewSigninResponseBody builds the HTTP response body from the result of the
// "signin" endpoint of the "secured_service" service.
func NewSigninResponseBody(res *securedservice.Creds) *SigninResponseBody {
	body := &SigninResponseBody{
		JWT:        res.JWT,
		APIKey:     res.APIKey,
		OauthToken: res.OauthToken,
	}
	return body
}

// NewSigninPayload builds a secured_service service signin endpoint payload.
func NewSigninPayload() *securedservice.SigninPayload {
	v := &securedservice.SigninPayload{}

	return v
}

// NewSecurePayload builds a secured_service service secure endpoint payload.
func NewSecurePayload(fail *bool, token string) *securedservice.SecurePayload {
	v := &securedservice.SecurePayload{}
	v.Fail = fail
	v.Token = token

	return v
}

// NewDoublySecurePayload builds a secured_service service doubly_secure
// endpoint payload.
func NewDoublySecurePayload(key string, token string) *securedservice.DoublySecurePayload {
	v := &securedservice.DoublySecurePayload{}
	v.Key = key
	v.Token = token

	return v
}

// NewAlsoDoublySecurePayload builds a secured_service service
// also_doubly_secure endpoint payload.
func NewAlsoDoublySecurePayload(key *string, oauthToken *string, token *string) *securedservice.AlsoDoublySecurePayload {
	v := &securedservice.AlsoDoublySecurePayload{}
	v.Key = key
	v.OauthToken = oauthToken
	v.Token = token

	return v
}
