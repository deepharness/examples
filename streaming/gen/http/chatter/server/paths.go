// Code generated by goa v3.7.6, DO NOT EDIT.
//
// HTTP request path constructors for the chatter service.
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package server

// LoginChatterPath returns the URL path to the chatter service login HTTP endpoint.
func LoginChatterPath() string {
	return "/login"
}

// EchoerChatterPath returns the URL path to the chatter service echoer HTTP endpoint.
func EchoerChatterPath() string {
	return "/echoer"
}

// ListenerChatterPath returns the URL path to the chatter service listener HTTP endpoint.
func ListenerChatterPath() string {
	return "/listener"
}

// SummaryChatterPath returns the URL path to the chatter service summary HTTP endpoint.
func SummaryChatterPath() string {
	return "/summary"
}

// SubscribeChatterPath returns the URL path to the chatter service subscribe HTTP endpoint.
func SubscribeChatterPath() string {
	return "/subscribe"
}

// HistoryChatterPath returns the URL path to the chatter service history HTTP endpoint.
func HistoryChatterPath() string {
	return "/history"
}
