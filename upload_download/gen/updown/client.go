// Code generated by goa v3.14.4, DO NOT EDIT.
//
// updown client
//
// Command:
// $ goa gen goa.design/examples/upload_download/design

package updown

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "updown" service client.
type Client struct {
	UploadEndpoint   goa.Endpoint
	DownloadEndpoint goa.Endpoint
}

// NewClient initializes a "updown" service client given the endpoints.
func NewClient(upload, download goa.Endpoint) *Client {
	return &Client{
		UploadEndpoint:   upload,
		DownloadEndpoint: download,
	}
}

// Upload calls the "upload" endpoint of the "updown" service.
// Upload may return the following errors:
//   - "invalid_media_type" (type *goa.ServiceError): Error returned when the Content-Type header does not define a multipart request.
//   - "invalid_multipart_request" (type *goa.ServiceError): Error returned when the request body is not a valid multipart content.
//   - "internal_error" (type *goa.ServiceError): Fault while processing upload.
//   - error: internal error
func (c *Client) Upload(ctx context.Context, p *UploadPayload, req io.ReadCloser) (err error) {
	_, err = c.UploadEndpoint(ctx, &UploadRequestData{Payload: p, Body: req})
	return
}

// Download calls the "download" endpoint of the "updown" service.
// Download may return the following errors:
//   - "invalid_file_path" (type *goa.ServiceError): Could not locate file for download
//   - "internal_error" (type *goa.ServiceError): Fault while processing download.
//   - error: internal error
func (c *Client) Download(ctx context.Context, p string) (res *DownloadResult, resp io.ReadCloser, err error) {
	var ires any
	ires, err = c.DownloadEndpoint(ctx, p)
	if err != nil {
		return
	}
	o := ires.(*DownloadResponseData)
	return o.Result, o.Body, nil
}
