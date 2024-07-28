// Code generated by goa v3.18.0, DO NOT EDIT.
//
// chatter views
//
// Command:
// $ goa gen goa.design/examples/streaming/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ChatSummaryCollection is the viewed result type that is projected based on a
// view.
type ChatSummaryCollection struct {
	// Type to project
	Projected ChatSummaryCollectionView
	// View to render
	View string
}

// ChatSummary is the viewed result type that is projected based on a view.
type ChatSummary struct {
	// Type to project
	Projected *ChatSummaryView
	// View to render
	View string
}

// EmptyView is a type that runs validations on a projected type.
type EmptyView struct {
}

// ChatSummaryCollectionView is a type that runs validations on a projected
// type.
type ChatSummaryCollectionView []*ChatSummaryView

// ChatSummaryView is a type that runs validations on a projected type.
type ChatSummaryView struct {
	// Message sent to the server
	Message *string
	// Length of the message sent
	Length *int
	// Time at which the message was sent
	SentAt *string
}

// EventView is a type that runs validations on a projected type.
type EventView struct {
	// Message sent to the server
	Message *string
	Action  *string
	// Time at which the message was added
	AddedAt *string
}

var (
	// ChatSummaryCollectionMap is a map indexing the attribute names of
	// ChatSummaryCollection by view name.
	ChatSummaryCollectionMap = map[string][]string{
		"tiny": {
			"message",
		},
		"default": {
			"message",
			"length",
			"sent_at",
		},
	}
	// ChatSummaryMap is a map indexing the attribute names of ChatSummary by view
	// name.
	ChatSummaryMap = map[string][]string{
		"tiny": {
			"message",
		},
		"default": {
			"message",
			"length",
			"sent_at",
		},
	}
)

// ValidateChatSummaryCollection runs the validations defined on the viewed
// result type ChatSummaryCollection.
func ValidateChatSummaryCollection(result ChatSummaryCollection) (err error) {
	switch result.View {
	case "tiny":
		err = ValidateChatSummaryCollectionViewTiny(result.Projected)
	case "default", "":
		err = ValidateChatSummaryCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"tiny", "default"})
	}
	return
}

// ValidateChatSummary runs the validations defined on the viewed result type
// ChatSummary.
func ValidateChatSummary(result *ChatSummary) (err error) {
	switch result.View {
	case "tiny":
		err = ValidateChatSummaryViewTiny(result.Projected)
	case "default", "":
		err = ValidateChatSummaryView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"tiny", "default"})
	}
	return
}

// ValidateEmptyView runs the validations defined on EmptyView.
func ValidateEmptyView(result *EmptyView) (err error) {

	return
}

// ValidateChatSummaryCollectionViewTiny runs the validations defined on
// ChatSummaryCollectionView using the "tiny" view.
func ValidateChatSummaryCollectionViewTiny(result ChatSummaryCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateChatSummaryViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateChatSummaryCollectionView runs the validations defined on
// ChatSummaryCollectionView using the "default" view.
func ValidateChatSummaryCollectionView(result ChatSummaryCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateChatSummaryView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateChatSummaryViewTiny runs the validations defined on ChatSummaryView
// using the "tiny" view.
func ValidateChatSummaryViewTiny(result *ChatSummaryView) (err error) {
	if result.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "result"))
	}
	return
}

// ValidateChatSummaryView runs the validations defined on ChatSummaryView
// using the "default" view.
func ValidateChatSummaryView(result *ChatSummaryView) (err error) {
	if result.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "result"))
	}
	if result.SentAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sent_at", "result"))
	}
	if result.SentAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.sent_at", *result.SentAt, goa.FormatDateTime))
	}
	return
}

// ValidateEventView runs the validations defined on EventView.
func ValidateEventView(result *EventView) (err error) {
	if result.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "result"))
	}
	if result.Action == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action", "result"))
	}
	if result.AddedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("added_at", "result"))
	}
	if result.Action != nil {
		if !(*result.Action == "added") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.action", *result.Action, []any{"added"}))
		}
	}
	if result.AddedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.added_at", *result.AddedAt, goa.FormatDateTime))
	}
	return
}
