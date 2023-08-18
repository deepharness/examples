// Code generated by goa v3.12.4, DO NOT EDIT.
//
// sommelier gRPC client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package client

import (
	"context"

	sommelierpb "goa.design/examples/cellar/gen/grpc/sommelier/pb"
	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildPickFunc builds the remote method to invoke for "sommelier" service
// "pick" endpoint.
func BuildPickFunc(grpccli sommelierpb.SommelierClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Pick(ctx, reqpb.(*sommelierpb.PickRequest), opts...)
		}
		return grpccli.Pick(ctx, &sommelierpb.PickRequest{}, opts...)
	}
}

// EncodePickRequest encodes requests sent to sommelier pick endpoint.
func EncodePickRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*sommelier.Criteria)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sommelier", "pick", "*sommelier.Criteria", v)
	}
	return NewProtoPickRequest(payload), nil
}

// DecodePickResponse decodes responses from the sommelier pick endpoint.
func DecodePickResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*sommelierpb.StoredBottleCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sommelier", "pick", "*sommelierpb.StoredBottleCollection", v)
	}
	res := NewPickResult(message)
	vres := sommelierviews.StoredBottleCollection{Projected: res, View: view}
	if err := sommelierviews.ValidateStoredBottleCollection(vres); err != nil {
		return nil, err
	}
	return sommelier.NewStoredBottleCollection(vres), nil
}
