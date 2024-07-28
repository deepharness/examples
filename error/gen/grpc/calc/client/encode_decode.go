// Code generated by goa v3.18.0, DO NOT EDIT.
//
// calc gRPC client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/error/design

package client

import (
	"context"

	calc "goa.design/examples/error/gen/calc"
	calcpb "goa.design/examples/error/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildDivideFunc builds the remote method to invoke for "calc" service
// "divide" endpoint.
func BuildDivideFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Divide(ctx, reqpb.(*calcpb.DivideRequest), opts...)
		}
		return grpccli.Divide(ctx, &calcpb.DivideRequest{}, opts...)
	}
}

// EncodeDivideRequest encodes requests sent to calc divide endpoint.
func EncodeDivideRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*calc.DividePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "divide", "*calc.DividePayload", v)
	}
	return NewProtoDivideRequest(payload), nil
}

// DecodeDivideResponse decodes responses from the calc divide endpoint.
func DecodeDivideResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*calcpb.DivideResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "divide", "*calcpb.DivideResponse", v)
	}
	res := NewDivideResult(message)
	return res, nil
}
