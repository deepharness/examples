// Code generated by goa v3.5.4, DO NOT EDIT.
//
// calc HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	calc "goa.design/examples/error/gen/calc"
	goahttp "goa.design/goa/v3/http"
)

// BuildDivideRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "divide" endpoint
func (c *Client) BuildDivideRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DivideCalcPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "divide", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDivideRequest returns an encoder for requests sent to the calc divide
// server.
func EncodeDivideRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*calc.DividePayload)
		if !ok {
			return goahttp.ErrInvalidType("calc", "divide", "*calc.DividePayload", v)
		}
		body := NewDivideRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("calc", "divide", err)
		}
		return nil
	}
}

// DecodeDivideResponse returns a decoder for responses returned by the calc
// divide endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDivideResponse may return the following errors:
//	- "div_by_zero" (type *calc.DivByZero): http.StatusBadRequest
//	- "timeout" (type *goa.ServiceError): http.StatusGatewayTimeout
//	- error: internal error
func DecodeDivideResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DivideResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "divide", err)
			}
			err = ValidateDivideResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("calc", "divide", err)
			}
			res := NewDivideResultOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body DivideDivByZeroResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "divide", err)
			}
			err = ValidateDivideDivByZeroResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("calc", "divide", err)
			}
			return nil, NewDivideDivByZero(&body)
		case http.StatusGatewayTimeout:
			var (
				body DivideTimeoutResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "divide", err)
			}
			err = ValidateDivideTimeoutResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("calc", "divide", err)
			}
			return nil, NewDivideTimeout(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "divide", resp.StatusCode, string(body))
		}
	}
}
