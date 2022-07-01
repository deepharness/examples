// Code generated by goa v3.7.7, DO NOT EDIT.
//
// chatter HTTP server
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package server

import (
	"context"
	"net/http"

	chatter "goa.design/examples/streaming/gen/chatter"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the chatter service endpoint HTTP handlers.
type Server struct {
	Mounts    []*MountPoint
	Login     http.Handler
	Echoer    http.Handler
	Listener  http.Handler
	Summary   http.Handler
	Subscribe http.Handler
	History   http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the chatter service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *chatter.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer *ConnConfigurer,
) *Server {
	if configurer == nil {
		configurer = &ConnConfigurer{}
	}
	return &Server{
		Mounts: []*MountPoint{
			{"Login", "POST", "/login"},
			{"Echoer", "GET", "/echoer"},
			{"Listener", "GET", "/listener"},
			{"Summary", "GET", "/summary"},
			{"Subscribe", "GET", "/subscribe"},
			{"History", "GET", "/history"},
		},
		Login:     NewLoginHandler(e.Login, mux, decoder, encoder, errhandler, formatter),
		Echoer:    NewEchoerHandler(e.Echoer, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.EchoerFn),
		Listener:  NewListenerHandler(e.Listener, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.ListenerFn),
		Summary:   NewSummaryHandler(e.Summary, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.SummaryFn),
		Subscribe: NewSubscribeHandler(e.Subscribe, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.SubscribeFn),
		History:   NewHistoryHandler(e.History, mux, decoder, encoder, errhandler, formatter, upgrader, configurer.HistoryFn),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "chatter" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Login = m(s.Login)
	s.Echoer = m(s.Echoer)
	s.Listener = m(s.Listener)
	s.Summary = m(s.Summary)
	s.Subscribe = m(s.Subscribe)
	s.History = m(s.History)
}

// Mount configures the mux to serve the chatter endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountLoginHandler(mux, h.Login)
	MountEchoerHandler(mux, h.Echoer)
	MountListenerHandler(mux, h.Listener)
	MountSummaryHandler(mux, h.Summary)
	MountSubscribeHandler(mux, h.Subscribe)
	MountHistoryHandler(mux, h.History)
}

// Mount configures the mux to serve the chatter endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountLoginHandler configures the mux to serve the "chatter" service "login"
// endpoint.
func MountLoginHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/login", f)
}

// NewLoginHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "login" endpoint.
func NewLoginHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeLoginRequest(mux, decoder)
		encodeResponse = EncodeLoginResponse(encoder)
		encodeError    = EncodeLoginError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "login")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountEchoerHandler configures the mux to serve the "chatter" service
// "echoer" endpoint.
func MountEchoerHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/echoer", f)
}

// NewEchoerHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "echoer" endpoint.
func NewEchoerHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeEchoerRequest(mux, decoder)
		encodeError   = EncodeEchoerError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "echoer")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &chatter.EchoerEndpointInput{
			Stream: &EchoerServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*chatter.EchoerPayload),
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}

// MountListenerHandler configures the mux to serve the "chatter" service
// "listener" endpoint.
func MountListenerHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/listener", f)
}

// NewListenerHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "listener" endpoint.
func NewListenerHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeListenerRequest(mux, decoder)
		encodeError   = EncodeListenerError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "listener")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &chatter.ListenerEndpointInput{
			Stream: &ListenerServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*chatter.ListenerPayload),
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}

// MountSummaryHandler configures the mux to serve the "chatter" service
// "summary" endpoint.
func MountSummaryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/summary", f)
}

// NewSummaryHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "summary" endpoint.
func NewSummaryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeSummaryRequest(mux, decoder)
		encodeError   = EncodeSummaryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "summary")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &chatter.SummaryEndpointInput{
			Stream: &SummaryServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*chatter.SummaryPayload),
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}

// MountSubscribeHandler configures the mux to serve the "chatter" service
// "subscribe" endpoint.
func MountSubscribeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/subscribe", f)
}

// NewSubscribeHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "subscribe" endpoint.
func NewSubscribeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeSubscribeRequest(mux, decoder)
		encodeError   = EncodeSubscribeError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "subscribe")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &chatter.SubscribeEndpointInput{
			Stream: &SubscribeServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*chatter.SubscribePayload),
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}

// MountHistoryHandler configures the mux to serve the "chatter" service
// "history" endpoint.
func MountHistoryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/history", f)
}

// NewHistoryHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "history" endpoint.
func NewHistoryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	upgrader goahttp.Upgrader,
	configurer goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeHistoryRequest(mux, decoder)
		encodeError   = EncodeHistoryError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "history")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		v := &chatter.HistoryEndpointInput{
			Stream: &HistoryServerStream{
				upgrader:   upgrader,
				configurer: configurer,
				cancel:     cancel,
				w:          w,
				r:          r,
			},
			Payload: payload.(*chatter.HistoryPayload),
		}
		_, err = endpoint(ctx, v)
		if err != nil {
			if _, werr := w.Write(nil); werr == http.ErrHijacked {
				// Response writer has been hijacked, do not encode the error
				errhandler(ctx, w, err)
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
	})
}
