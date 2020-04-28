// Code generated by goa v3.1.1, DO NOT EDIT.
//
// tus HTTP server types
//
// Command:
// $ goa gen goa.design/examples/tus/design -o
// $(GOPATH)/src/goa.design/examples/tus

package server

import (
	tus "goa.design/examples/tus/gen/tus"
	goa "goa.design/goa/v3/pkg"
)

// HeadNotFoundResponseBody is the type of the "tus" service "head" endpoint
// HTTP response body for the "NotFound" error.
type HeadNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// HeadGoneResponseBody is the type of the "tus" service "head" endpoint HTTP
// response body for the "Gone" error.
type HeadGoneResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PatchInvalidContentTypeResponseBody is the type of the "tus" service "patch"
// endpoint HTTP response body for the "InvalidContentType" error.
type PatchInvalidContentTypeResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PatchInvalidOffsetResponseBody is the type of the "tus" service "patch"
// endpoint HTTP response body for the "InvalidOffset" error.
type PatchInvalidOffsetResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PatchNotFoundResponseBody is the type of the "tus" service "patch" endpoint
// HTTP response body for the "NotFound" error.
type PatchNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PatchGoneResponseBody is the type of the "tus" service "patch" endpoint HTTP
// response body for the "Gone" error.
type PatchGoneResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PatchInvalidChecksumAlgorithmResponseBody is the type of the "tus" service
// "patch" endpoint HTTP response body for the "InvalidChecksumAlgorithm" error.
type PatchInvalidChecksumAlgorithmResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PatchChecksumMismatchResponseBody is the type of the "tus" service "patch"
// endpoint HTTP response body for the "ChecksumMismatch" error.
type PatchChecksumMismatchResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PatchInternalResponseBody is the type of the "tus" service "patch" endpoint
// HTTP response body for the "Internal" error.
type PatchInternalResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PostMissingHeaderResponseBody is the type of the "tus" service "post"
// endpoint HTTP response body for the "MissingHeader" error.
type PostMissingHeaderResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PostInvalidDeferLengthResponseBody is the type of the "tus" service "post"
// endpoint HTTP response body for the "InvalidDeferLength" error.
type PostInvalidDeferLengthResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PostInvalidChecksumAlgorithmResponseBody is the type of the "tus" service
// "post" endpoint HTTP response body for the "InvalidChecksumAlgorithm" error.
type PostInvalidChecksumAlgorithmResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PostMaximumSizeExceededResponseBody is the type of the "tus" service "post"
// endpoint HTTP response body for the "MaximumSizeExceeded" error.
type PostMaximumSizeExceededResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// PostChecksumMismatchResponseBody is the type of the "tus" service "post"
// endpoint HTTP response body for the "ChecksumMismatch" error.
type PostChecksumMismatchResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteNotFoundResponseBody is the type of the "tus" service "delete"
// endpoint HTTP response body for the "NotFound" error.
type DeleteNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteGoneResponseBody is the type of the "tus" service "delete" endpoint
// HTTP response body for the "Gone" error.
type DeleteGoneResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewHeadNotFoundResponseBody builds the HTTP response body from the result of
// the "head" endpoint of the "tus" service.
func NewHeadNotFoundResponseBody(res *goa.ServiceError) *HeadNotFoundResponseBody {
	body := &HeadNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewHeadGoneResponseBody builds the HTTP response body from the result of the
// "head" endpoint of the "tus" service.
func NewHeadGoneResponseBody(res *goa.ServiceError) *HeadGoneResponseBody {
	body := &HeadGoneResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPatchInvalidContentTypeResponseBody builds the HTTP response body from
// the result of the "patch" endpoint of the "tus" service.
func NewPatchInvalidContentTypeResponseBody(res *goa.ServiceError) *PatchInvalidContentTypeResponseBody {
	body := &PatchInvalidContentTypeResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPatchInvalidOffsetResponseBody builds the HTTP response body from the
// result of the "patch" endpoint of the "tus" service.
func NewPatchInvalidOffsetResponseBody(res *goa.ServiceError) *PatchInvalidOffsetResponseBody {
	body := &PatchInvalidOffsetResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPatchNotFoundResponseBody builds the HTTP response body from the result
// of the "patch" endpoint of the "tus" service.
func NewPatchNotFoundResponseBody(res *goa.ServiceError) *PatchNotFoundResponseBody {
	body := &PatchNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPatchGoneResponseBody builds the HTTP response body from the result of
// the "patch" endpoint of the "tus" service.
func NewPatchGoneResponseBody(res *goa.ServiceError) *PatchGoneResponseBody {
	body := &PatchGoneResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPatchInvalidChecksumAlgorithmResponseBody builds the HTTP response body
// from the result of the "patch" endpoint of the "tus" service.
func NewPatchInvalidChecksumAlgorithmResponseBody(res *goa.ServiceError) *PatchInvalidChecksumAlgorithmResponseBody {
	body := &PatchInvalidChecksumAlgorithmResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPatchChecksumMismatchResponseBody builds the HTTP response body from the
// result of the "patch" endpoint of the "tus" service.
func NewPatchChecksumMismatchResponseBody(res *goa.ServiceError) *PatchChecksumMismatchResponseBody {
	body := &PatchChecksumMismatchResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPatchInternalResponseBody builds the HTTP response body from the result
// of the "patch" endpoint of the "tus" service.
func NewPatchInternalResponseBody(res *goa.ServiceError) *PatchInternalResponseBody {
	body := &PatchInternalResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPostMissingHeaderResponseBody builds the HTTP response body from the
// result of the "post" endpoint of the "tus" service.
func NewPostMissingHeaderResponseBody(res *goa.ServiceError) *PostMissingHeaderResponseBody {
	body := &PostMissingHeaderResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPostInvalidDeferLengthResponseBody builds the HTTP response body from the
// result of the "post" endpoint of the "tus" service.
func NewPostInvalidDeferLengthResponseBody(res *goa.ServiceError) *PostInvalidDeferLengthResponseBody {
	body := &PostInvalidDeferLengthResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPostInvalidChecksumAlgorithmResponseBody builds the HTTP response body
// from the result of the "post" endpoint of the "tus" service.
func NewPostInvalidChecksumAlgorithmResponseBody(res *goa.ServiceError) *PostInvalidChecksumAlgorithmResponseBody {
	body := &PostInvalidChecksumAlgorithmResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPostMaximumSizeExceededResponseBody builds the HTTP response body from
// the result of the "post" endpoint of the "tus" service.
func NewPostMaximumSizeExceededResponseBody(res *goa.ServiceError) *PostMaximumSizeExceededResponseBody {
	body := &PostMaximumSizeExceededResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPostChecksumMismatchResponseBody builds the HTTP response body from the
// result of the "post" endpoint of the "tus" service.
func NewPostChecksumMismatchResponseBody(res *goa.ServiceError) *PostChecksumMismatchResponseBody {
	body := &PostChecksumMismatchResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "tus" service.
func NewDeleteNotFoundResponseBody(res *goa.ServiceError) *DeleteNotFoundResponseBody {
	body := &DeleteNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteGoneResponseBody builds the HTTP response body from the result of
// the "delete" endpoint of the "tus" service.
func NewDeleteGoneResponseBody(res *goa.ServiceError) *DeleteGoneResponseBody {
	body := &DeleteGoneResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewHeadPayload builds a tus service head endpoint payload.
func NewHeadPayload(id string, tusResumable string) *tus.HeadPayload {
	v := &tus.HeadPayload{}
	v.ID = id
	v.TusResumable = tusResumable

	return v
}

// NewPatchPayload builds a tus service patch endpoint payload.
func NewPatchPayload(id string, tusResumable string, uploadOffset int64, uploadChecksum *string) *tus.PatchPayload {
	v := &tus.PatchPayload{}
	v.ID = id
	v.TusResumable = tusResumable
	v.UploadOffset = uploadOffset
	v.UploadChecksum = uploadChecksum

	return v
}

// NewPostPayload builds a tus service post endpoint payload.
func NewPostPayload(tusResumable string, uploadLength *int64, uploadDeferLength *int, uploadChecksum *string, uploadMetadata *string, tusMaxSize *int64) *tus.PostPayload {
	v := &tus.PostPayload{}
	v.TusResumable = tusResumable
	v.UploadLength = uploadLength
	v.UploadDeferLength = uploadDeferLength
	v.UploadChecksum = uploadChecksum
	v.UploadMetadata = uploadMetadata
	v.TusMaxSize = tusMaxSize

	return v
}

// NewDeletePayload builds a tus service delete endpoint payload.
func NewDeletePayload(id string, tusResumable string) *tus.DeletePayload {
	v := &tus.DeletePayload{}
	v.ID = id
	v.TusResumable = tusResumable

	return v
}
