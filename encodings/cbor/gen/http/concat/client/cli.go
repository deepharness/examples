// Code generated by goa v3.16.2, DO NOT EDIT.
//
// concat HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/encodings/cbor/design

package client

import (
	concat "goa.design/examples/encodings/cbor/gen/concat"
)

// BuildConcatPayload builds the payload for the concat concat endpoint from
// CLI flags.
func BuildConcatPayload(concatConcatA string, concatConcatB string) (*concat.ConcatPayload, error) {
	var a string
	{
		a = concatConcatA
	}
	var b string
	{
		b = concatConcatB
	}
	v := &concat.ConcatPayload{}
	v.A = a
	v.B = b

	return v, nil
}
