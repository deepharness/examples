// Code generated by goa v3.16.1, DO NOT EDIT.
//
// updown service
//
// Command:
// $ goa gen goa.design/examples/upload_download/design

package updown

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Service updown demonstrates how to implement upload and download of files in
// Goa without having to load the entire content in memory first.
// The upload method uses SkipRequestBodyEncodeDecode to delegate reading the
// HTTP
// request body to the service logic. This alleviates the need for loading the
// full body content in memory first to decode it into a data structure. Note
// that
// using SkipRequestBodyDecode is incompatible with gRPC and can only be used on
// methods that only define a HTTP transport mapping. This example
// implementation
// leverages package "mime/multipart" to read the request body.
// Similarly the download method uses SkipResponseBodyEncodeDecode to stream the
// file to the client without requiring to load the complete content in memory
// first. As with SkipRequestBodyDecode using SkipResponseBodyEncodeDecode is
// incompatible with gRPC.
type Service interface {
	// Upload implements upload.
	Upload(context.Context, *UploadPayload, io.ReadCloser) (err error)
	// Download implements download.
	Download(context.Context, string) (res *DownloadResult, body io.ReadCloser, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "upload_download"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "updown"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"upload", "download"}

// DownloadResult is the result type of the updown service download method.
type DownloadResult struct {
	// Length is the downloaded content length in bytes.
	Length int64
}

// UploadPayload is the payload type of the updown service upload method.
type UploadPayload struct {
	// Content-Type header, must define value for multipart boundary.
	ContentType string
	// Dir is the relative path to the file directory where the uploaded content is
	// saved.
	Dir string
}

// MakeInvalidMediaType builds a goa.ServiceError from an error.
func MakeInvalidMediaType(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid_media_type", false, false, false)
}

// MakeInvalidMultipartRequest builds a goa.ServiceError from an error.
func MakeInvalidMultipartRequest(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid_multipart_request", false, false, false)
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "internal_error", false, false, false)
}

// MakeInvalidFilePath builds a goa.ServiceError from an error.
func MakeInvalidFilePath(err error) *goa.ServiceError {
	return goa.NewServiceError(err, "invalid_file_path", false, false, false)
}
