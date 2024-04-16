// Code generated by goa v3.16.1, DO NOT EDIT.
//
// calc gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/error/design

package client

import (
	"encoding/json"
	"fmt"

	calc "goa.design/examples/error/gen/calc"
	calcpb "goa.design/examples/error/gen/grpc/calc/pb"
)

// BuildDividePayload builds the payload for the calc divide endpoint from CLI
// flags.
func BuildDividePayload(calcDivideMessage string) (*calc.DividePayload, error) {
	var err error
	var message calcpb.DivideRequest
	{
		if calcDivideMessage != "" {
			err = json.Unmarshal([]byte(calcDivideMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"dividend\": 6001261670528037577,\n      \"divisor\": 5850431520333673251\n   }'")
			}
		}
	}
	v := &calc.DividePayload{
		Dividend: int(message.Dividend),
		Divisor:  int(message.Divisor),
	}

	return v, nil
}
