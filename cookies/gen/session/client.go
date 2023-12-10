// Code generated by goa v3.14.1, DO NOT EDIT.
//
// session client
//
// Command:
// $ goa gen goa.design/examples/cookies/design

package session

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "session" service client.
type Client struct {
	CreateSessionEndpoint goa.Endpoint
	UseSessionEndpoint    goa.Endpoint
}

// NewClient initializes a "session" service client given the endpoints.
func NewClient(createSession, useSession goa.Endpoint) *Client {
	return &Client{
		CreateSessionEndpoint: createSession,
		UseSessionEndpoint:    useSession,
	}
}

// CreateSession calls the "create_session" endpoint of the "session" service.
func (c *Client) CreateSession(ctx context.Context, p *CreateSessionPayload) (res *CreateSessionResult, err error) {
	var ires any
	ires, err = c.CreateSessionEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CreateSessionResult), nil
}

// UseSession calls the "use_session" endpoint of the "session" service.
func (c *Client) UseSession(ctx context.Context, p *UseSessionPayload) (res *UseSessionResult, err error) {
	var ires any
	ires, err = c.UseSessionEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UseSessionResult), nil
}
