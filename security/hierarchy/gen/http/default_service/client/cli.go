// Code generated by goa v3.16.1, DO NOT EDIT.
//
// default_service HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design

package client

import (
	defaultservice "goa.design/examples/security/hierarchy/gen/default_service"
)

// BuildDefaultPayload builds the payload for the default_service default
// endpoint from CLI flags.
func BuildDefaultPayload(defaultServiceDefaultUsername string, defaultServiceDefaultPassword string) (*defaultservice.DefaultPayload, error) {
	var username string
	{
		username = defaultServiceDefaultUsername
	}
	var password string
	{
		password = defaultServiceDefaultPassword
	}
	v := &defaultservice.DefaultPayload{}
	v.Username = username
	v.Password = password

	return v, nil
}
