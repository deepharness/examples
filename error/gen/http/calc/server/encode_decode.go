// Code generated by goa v3.2.5, DO NOT EDIT.
//
// calc HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package server

import (
	"context"
	"io"
	"net/http"

	calc "goa.design/examples/error/gen/calc"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeDivideResponse returns an encoder for responses returned by the calc
// divide endpoint.
func EncodeDivideResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*calc.DivideResult)
		enc := encoder(ctx, w)
		body := NewDivideResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDivideRequest returns a decoder for requests sent to the calc divide
// endpoint.
func DecodeDivideRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body DivideRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateDivideRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewDividePayload(&body)

		return payload, nil
	}
}

// EncodeDivideError returns an encoder for errors returned by the divide calc
// endpoint.
func EncodeDivideError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "div_by_zero":
			res := v.(*calc.DivByZero)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDivideDivByZeroResponseBody(res)
			}
			w.Header().Set("goa-error", "div_by_zero")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "timeout":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDivideTimeoutResponseBody(res)
			}
			w.Header().Set("goa-error", "timeout")
			w.WriteHeader(http.StatusGatewayTimeout)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
