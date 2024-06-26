// Code generated by goa v3.15.2, DO NOT EDIT.
//
// hello views
//
// Command:
// $ goa gen goa.design/examples/httpstatus/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Hello is the viewed result type that is projected based on a view.
type Hello struct {
	// Type to project
	Projected *HelloView
	// View to render
	View string
}

// HelloView is a type that runs validations on a projected type.
type HelloView struct {
	// The greeting message
	Greeting *string
	Outcome  *string `json:"-"`
}

var (
	// HelloMap is a map indexing the attribute names of Hello by view name.
	HelloMap = map[string][]string{
		"default": {
			"greeting",
			"outcome",
		},
	}
)

// ValidateHello runs the validations defined on the viewed result type Hello.
func ValidateHello(result *Hello) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateHelloView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateHelloView runs the validations defined on HelloView using the
// "default" view.
func ValidateHelloView(result *HelloView) (err error) {
	if result.Outcome == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("outcome", "result"))
	}
	if result.Greeting == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("greeting", "result"))
	}
	return
}
