// Code generated by goa v3.11.2, DO NOT EDIT.
//
// session client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/cookies/design -o cookies

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the session service endpoint HTTP clients.
type Client struct {
	// CreateSession Doer is the HTTP client used to make requests to the
	// create_session endpoint.
	CreateSessionDoer goahttp.Doer

	// UseSession Doer is the HTTP client used to make requests to the use_session
	// endpoint.
	UseSessionDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the session service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CreateSessionDoer:   doer,
		UseSessionDoer:      doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// CreateSession returns an endpoint that makes HTTP requests to the session
// service create_session server.
func (c *Client) CreateSession() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateSessionRequest(c.encoder)
		decodeResponse = DecodeCreateSessionResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateSessionRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateSessionDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("session", "create_session", err)
		}
		return decodeResponse(resp)
	}
}

// UseSession returns an endpoint that makes HTTP requests to the session
// service use_session server.
func (c *Client) UseSession() goa.Endpoint {
	var (
		encodeRequest  = EncodeUseSessionRequest(c.encoder)
		decodeResponse = DecodeUseSessionResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUseSessionRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UseSessionDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("session", "use_session", err)
		}
		return decodeResponse(resp)
	}
}
