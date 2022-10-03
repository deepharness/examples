// Code generated by goa v3.9.1, DO NOT EDIT.
//
// api_key_service HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design -o security/hierarchy

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	apikeyservice "goa.design/examples/security/hierarchy/gen/api_key_service"
	goahttp "goa.design/goa/v3/http"
)

// BuildDefaultRequest instantiates a HTTP request object with method and path
// set to call the "api_key_service" service "default" endpoint
func (c *Client) BuildDefaultRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DefaultAPIKeyServicePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("api_key_service", "default", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDefaultRequest returns an encoder for requests sent to the
// api_key_service default server.
func EncodeDefaultRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*apikeyservice.DefaultPayload)
		if !ok {
			return goahttp.ErrInvalidType("api_key_service", "default", "*apikeyservice.DefaultPayload", v)
		}
		{
			head := p.Key
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeDefaultResponse returns a decoder for responses returned by the
// api_key_service default endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeDefaultResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("api_key_service", "default", resp.StatusCode, string(body))
		}
	}
}

// BuildSecureRequest instantiates a HTTP request object with method and path
// set to call the "api_key_service" service "secure" endpoint
func (c *Client) BuildSecureRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SecureAPIKeyServicePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("api_key_service", "secure", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSecureRequest returns an encoder for requests sent to the
// api_key_service secure server.
func EncodeSecureRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*apikeyservice.SecurePayload)
		if !ok {
			return goahttp.ErrInvalidType("api_key_service", "secure", "*apikeyservice.SecurePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		return nil
	}
}

// DecodeSecureResponse returns a decoder for responses returned by the
// api_key_service secure endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeSecureResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("api_key_service", "secure", resp.StatusCode, string(body))
		}
	}
}
