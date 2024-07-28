// Code generated by goa v3.18.0, DO NOT EDIT.
//
// sommelier HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package server

import (
	"context"
	"errors"
	"io"
	"net/http"

	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodePickResponse returns an encoder for responses returned by the
// sommelier pick endpoint.
func EncodePickResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res := v.(sommelierviews.StoredBottleCollection)
		enc := encoder(ctx, w)
		body := NewStoredBottleResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodePickRequest returns a decoder for requests sent to the sommelier pick
// endpoint.
func DecodePickRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			body PickRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			var gerr *goa.ServiceError
			if errors.As(err, &gerr) {
				return nil, gerr
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewPickCriteria(&body)

		return payload, nil
	}
}

// EncodePickError returns an encoder for errors returned by the pick sommelier
// endpoint.
func EncodePickError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "no_criteria":
			var res sommelier.NoCriteria
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "no_match":
			var res sommelier.NoMatch
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalSommelierviewsStoredBottleViewToStoredBottleResponse builds a value
// of type *StoredBottleResponse from a value of type
// *sommelierviews.StoredBottleView.
func marshalSommelierviewsStoredBottleViewToStoredBottleResponse(v *sommelierviews.StoredBottleView) *StoredBottleResponse {
	res := &StoredBottleResponse{
		ID:          *v.ID,
		Name:        *v.Name,
		Vintage:     *v.Vintage,
		Description: v.Description,
		Rating:      v.Rating,
	}
	if v.Winery != nil {
		res.Winery = marshalSommelierviewsWineryViewToWineryResponseTiny(v.Winery)
	}
	if v.Composition != nil {
		res.Composition = make([]*ComponentResponse, len(v.Composition))
		for i, val := range v.Composition {
			res.Composition[i] = marshalSommelierviewsComponentViewToComponentResponse(val)
		}
	}

	return res
}

// marshalSommelierviewsWineryViewToWineryResponseTiny builds a value of type
// *WineryResponseTiny from a value of type *sommelierviews.WineryView.
func marshalSommelierviewsWineryViewToWineryResponseTiny(v *sommelierviews.WineryView) *WineryResponseTiny {
	res := &WineryResponseTiny{
		Name: *v.Name,
	}

	return res
}

// marshalSommelierviewsComponentViewToComponentResponse builds a value of type
// *ComponentResponse from a value of type *sommelierviews.ComponentView.
func marshalSommelierviewsComponentViewToComponentResponse(v *sommelierviews.ComponentView) *ComponentResponse {
	if v == nil {
		return nil
	}
	res := &ComponentResponse{
		Varietal:   *v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}
