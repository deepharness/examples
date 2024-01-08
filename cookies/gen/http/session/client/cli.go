// Code generated by goa v3.14.4, DO NOT EDIT.
//
// session HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/cookies/design

package client

import (
	"encoding/json"
	"fmt"

	session "goa.design/examples/cookies/gen/session"
)

// BuildCreateSessionPayload builds the payload for the session create_session
// endpoint from CLI flags.
func BuildCreateSessionPayload(sessionCreateSessionBody string) (*session.CreateSessionPayload, error) {
	var err error
	var body CreateSessionRequestBody
	{
		err = json.Unmarshal([]byte(sessionCreateSessionBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Ut quas sit at illum tempore ad.\"\n   }'")
		}
	}
	v := &session.CreateSessionPayload{
		Name: body.Name,
	}

	return v, nil
}

// BuildUseSessionPayload builds the payload for the session use_session
// endpoint from CLI flags.
func BuildUseSessionPayload(sessionUseSessionSessionID string) (*session.UseSessionPayload, error) {
	var sessionID string
	{
		sessionID = sessionUseSessionSessionID
	}
	v := &session.UseSessionPayload{}
	v.SessionID = sessionID

	return v, nil
}
