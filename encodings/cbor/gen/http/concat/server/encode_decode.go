// Code generated by goa v3.12.4, DO NOT EDIT.
//
// concat HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/encodings/cbor/design -o encodings/cbor

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeConcatResponse returns an encoder for responses returned by the concat
// concat endpoint.
func EncodeConcatResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeConcatRequest returns a decoder for requests sent to the concat concat
// endpoint.
func DecodeConcatRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			a string
			b string

			params = mux.Vars(r)
		)
		a = params["a"]
		b = params["b"]
		payload := NewConcatPayload(a, b)

		return payload, nil
	}
}
