// Code generated by goa v3.16.1, DO NOT EDIT.
//
// tus client
//
// Command:
// $ goa gen goa.design/examples/tus/design

package tus

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "tus" service client.
type Client struct {
	HeadEndpoint    goa.Endpoint
	PatchEndpoint   goa.Endpoint
	OptionsEndpoint goa.Endpoint
	PostEndpoint    goa.Endpoint
	DeleteEndpoint  goa.Endpoint
}

// NewClient initializes a "tus" service client given the endpoints.
func NewClient(head, patch, options, post, delete_ goa.Endpoint) *Client {
	return &Client{
		HeadEndpoint:    head,
		PatchEndpoint:   patch,
		OptionsEndpoint: options,
		PostEndpoint:    post,
		DeleteEndpoint:  delete_,
	}
}

// Head calls the "head" endpoint of the "tus" service.
// Head may return the following errors:
//   - "NotFound" (type *goa.ServiceError): If the resource is not found, the Server SHOULD return either the 404 Not Found, 410 Gone or 403 Forbidden status without the Upload-Offset header.
//   - "Gone" (type *goa.ServiceError): If the resource is not found, the Server SHOULD return either the 404 Not Found, 410 Gone or 403 Forbidden status without the Upload-Offset header.
//   - "InvalidTusResumable" (type *ErrInvalidTUSResumable): If the version specified by the Client is not supported by the Server, it MUST respond with the 412 Precondition Failed status.
//   - error: internal error
func (c *Client) Head(ctx context.Context, p *HeadPayload) (res *HeadResult, err error) {
	var ires any
	ires, err = c.HeadEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*HeadResult), nil
}

// Patch calls the "patch" endpoint of the "tus" service.
// Patch may return the following errors:
//   - "InvalidContentType" (type *goa.ServiceError): All PATCH requests MUST use Content-Type: application/offset+octet-stream, otherwise the server SHOULD return a 415 Unsupported Media Type status.
//   - "InvalidOffset" (type *goa.ServiceError): If the offsets do not match, the Server MUST respond with the 409 Conflict status without modifying the upload resource.
//   - "NotFound" (type *goa.ServiceError): If a Client does attempt to resume an upload which has since been removed by the Server, the Server SHOULD respond with the404 Not Found or 410 Gone status.
//   - "Gone" (type *goa.ServiceError): If a Client does attempt to resume an upload which has since been removed by the Server, the Server SHOULD respond with the404 Not Found or 410 Gone status.
//   - "InvalidChecksumAlgorithm" (type *goa.ServiceError): The checksum algorithm is not supported by the server.
//   - "ChecksumMismatch" (type *goa.ServiceError): The checksums mismatch.
//   - "Internal" (type *goa.ServiceError): Internal error
//   - "InvalidTusResumable" (type *ErrInvalidTUSResumable): If the version specified by the Client is not supported by the Server, it MUST respond with the 412 Precondition Failed status.
//   - error: internal error
func (c *Client) Patch(ctx context.Context, p *PatchPayload, req io.ReadCloser) (res *PatchResult, err error) {
	var ires any
	ires, err = c.PatchEndpoint(ctx, &PatchRequestData{Payload: p, Body: req})
	if err != nil {
		return
	}
	return ires.(*PatchResult), nil
}

// Options calls the "options" endpoint of the "tus" service.
// Options may return the following errors:
//   - "InvalidTusResumable" (type *ErrInvalidTUSResumable): If the version specified by the Client is not supported by the Server, it MUST respond with the 412 Precondition Failed status.
//   - error: internal error
func (c *Client) Options(ctx context.Context) (res *OptionsResult, err error) {
	var ires any
	ires, err = c.OptionsEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*OptionsResult), nil
}

// Post calls the "post" endpoint of the "tus" service.
// Post may return the following errors:
//   - "MissingHeader" (type *goa.ServiceError): The request MUST include one of the following headers: a) Upload-Length -or- b) Upload-Defer-Length: 1
//   - "InvalidDeferLength" (type *goa.ServiceError): If the Upload-Defer-Length header contains any other value than 1 the server should return a 400 Bad Request status.
//   - "MaximumSizeExceeded" (type *goa.ServiceError): If the length of the upload exceeds the maximum, which MAY be specified using the Tus-Max-Size header, the Server MUST respond with the 413 Request Entity Too Large status.
//   - "InvalidChecksumAlgorithm" (type *goa.ServiceError): The checksum algorithm is not supported by the server.
//   - "ChecksumMismatch" (type *goa.ServiceError): The checksums mismatch.
//   - "InvalidTusResumable" (type *ErrInvalidTUSResumable): If the version specified by the Client is not supported by the Server, it MUST respond with the 412 Precondition Failed status.
//   - error: internal error
func (c *Client) Post(ctx context.Context, p *PostPayload, req io.ReadCloser) (res *PostResult, err error) {
	var ires any
	ires, err = c.PostEndpoint(ctx, &PostRequestData{Payload: p, Body: req})
	if err != nil {
		return
	}
	return ires.(*PostResult), nil
}

// Delete calls the "delete" endpoint of the "tus" service.
// Delete may return the following errors:
//   - "NotFound" (type *goa.ServiceError): For all future requests to this URL, the Server SHOULD respond with the 404 Not Found or 410 Gone status.
//   - "Gone" (type *goa.ServiceError): For all future requests to this URL, the Server SHOULD respond with the 404 Not Found or 410 Gone status.
//   - "InvalidTusResumable" (type *ErrInvalidTUSResumable): If the version specified by the Client is not supported by the Server, it MUST respond with the 412 Precondition Failed status.
//   - error: internal error
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (res *DeleteResult, err error) {
	var ires any
	ires, err = c.DeleteEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DeleteResult), nil
}
