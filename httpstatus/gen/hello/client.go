// Code generated by goa v3.15.2, DO NOT EDIT.
//
// hello client
//
// Command:
// $ goa gen goa.design/examples/httpstatus/design

package hello

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "hello" service client.
type Client struct {
	HelloEndpointEndpoint goa.Endpoint
}

// NewClient initializes a "hello" service client given the endpoints.
func NewClient(helloEndpoint goa.Endpoint) *Client {
	return &Client{
		HelloEndpointEndpoint: helloEndpoint,
	}
}

// HelloEndpoint calls the "hello" endpoint of the "hello" service.
func (c *Client) HelloEndpoint(ctx context.Context, p *HelloPayload) (res *Hello, err error) {
	var ires any
	ires, err = c.HelloEndpointEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Hello), nil
}