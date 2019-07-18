// Code generated by goa v3.0.3, DO NOT EDIT.
//
// chatter HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o
// $(GOPATH)/src/goa.design/examples/streaming

package client

import (
	chatter "goa.design/examples/streaming/gen/chatter"
)

// BuildLoginPayload builds the payload for the chatter login endpoint from CLI
// flags.
func BuildLoginPayload(chatterLoginUser string, chatterLoginPassword string) (*chatter.LoginPayload, error) {
	var user string
	{
		user = chatterLoginUser
	}
	var password string
	{
		password = chatterLoginPassword
	}
	payload := &chatter.LoginPayload{
		User:     user,
		Password: password,
	}
	return payload, nil
}

// BuildEchoerPayload builds the payload for the chatter echoer endpoint from
// CLI flags.
func BuildEchoerPayload(chatterEchoerToken string) (*chatter.EchoerPayload, error) {
	var token string
	{
		token = chatterEchoerToken
	}
	payload := &chatter.EchoerPayload{
		Token: token,
	}
	return payload, nil
}

// BuildListenerPayload builds the payload for the chatter listener endpoint
// from CLI flags.
func BuildListenerPayload(chatterListenerToken string) (*chatter.ListenerPayload, error) {
	var token string
	{
		token = chatterListenerToken
	}
	payload := &chatter.ListenerPayload{
		Token: token,
	}
	return payload, nil
}

// BuildSummaryPayload builds the payload for the chatter summary endpoint from
// CLI flags.
func BuildSummaryPayload(chatterSummaryToken string) (*chatter.SummaryPayload, error) {
	var token string
	{
		token = chatterSummaryToken
	}
	payload := &chatter.SummaryPayload{
		Token: token,
	}
	return payload, nil
}

// BuildSubscribePayload builds the payload for the chatter subscribe endpoint
// from CLI flags.
func BuildSubscribePayload(chatterSubscribeToken string) (*chatter.SubscribePayload, error) {
	var token string
	{
		token = chatterSubscribeToken
	}
	payload := &chatter.SubscribePayload{
		Token: token,
	}
	return payload, nil
}

// BuildHistoryPayload builds the payload for the chatter history endpoint from
// CLI flags.
func BuildHistoryPayload(chatterHistoryView string, chatterHistoryToken string) (*chatter.HistoryPayload, error) {
	var view *string
	{
		if chatterHistoryView != "" {
			view = &chatterHistoryView
		}
	}
	var token string
	{
		token = chatterHistoryToken
	}
	payload := &chatter.HistoryPayload{
		View:  view,
		Token: token,
	}
	return payload, nil
}
