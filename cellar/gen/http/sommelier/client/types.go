// Code generated by goa v3.7.1, DO NOT EDIT.
//
// sommelier HTTP client types
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package client

import (
	"unicode/utf8"

	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
	goa "goa.design/goa/v3/pkg"
)

// PickRequestBody is the type of the "sommelier" service "pick" endpoint HTTP
// request body.
type PickRequestBody struct {
	// Name of bottle to pick
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Varietals in preference order
	Varietal []string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Winery of bottle to pick
	Winery *string `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
}

// PickResponseBody is the type of the "sommelier" service "pick" endpoint HTTP
// response body.
type PickResponseBody []*StoredBottleResponse

// StoredBottleResponse is used to define fields on response body types.
type StoredBottleResponse struct {
	// ID is the unique id of the bottle.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of bottle
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Winery that produces wine
	Winery *WineryResponse `form:"winery,omitempty" json:"winery,omitempty" xml:"winery,omitempty"`
	// Vintage of bottle
	Vintage *uint32 `form:"vintage,omitempty" json:"vintage,omitempty" xml:"vintage,omitempty"`
	// Composition is the list of grape varietals and associated percentage.
	Composition []*ComponentResponse `form:"composition,omitempty" json:"composition,omitempty" xml:"composition,omitempty"`
	// Description of bottle
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32 `form:"rating,omitempty" json:"rating,omitempty" xml:"rating,omitempty"`
}

// WineryResponse is used to define fields on response body types.
type WineryResponse struct {
	// Name of winery
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Region of winery
	Region *string `form:"region,omitempty" json:"region,omitempty" xml:"region,omitempty"`
	// Country of winery
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Winery website URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// ComponentResponse is used to define fields on response body types.
type ComponentResponse struct {
	// Grape varietal
	Varietal *string `form:"varietal,omitempty" json:"varietal,omitempty" xml:"varietal,omitempty"`
	// Percentage of varietal in wine
	Percentage *uint32 `form:"percentage,omitempty" json:"percentage,omitempty" xml:"percentage,omitempty"`
}

// NewPickRequestBody builds the HTTP request body from the payload of the
// "pick" endpoint of the "sommelier" service.
func NewPickRequestBody(p *sommelier.Criteria) *PickRequestBody {
	body := &PickRequestBody{
		Name:   p.Name,
		Winery: p.Winery,
	}
	if p.Varietal != nil {
		body.Varietal = make([]string, len(p.Varietal))
		for i, val := range p.Varietal {
			body.Varietal[i] = val
		}
	}
	return body
}

// NewPickStoredBottleCollectionOK builds a "sommelier" service "pick" endpoint
// result from a HTTP "OK" response.
func NewPickStoredBottleCollectionOK(body PickResponseBody) sommelierviews.StoredBottleCollectionView {
	v := make([]*sommelierviews.StoredBottleView, len(body))
	for i, val := range body {
		v[i] = unmarshalStoredBottleResponseToSommelierviewsStoredBottleView(val)
	}

	return v
}

// NewPickNoCriteria builds a sommelier service pick endpoint no_criteria error.
func NewPickNoCriteria(body string) sommelier.NoCriteria {
	v := sommelier.NoCriteria(body)

	return v
}

// NewPickNoMatch builds a sommelier service pick endpoint no_match error.
func NewPickNoMatch(body string) sommelier.NoMatch {
	v := sommelier.NoMatch(body)

	return v
}

// ValidateStoredBottleResponse runs the validations defined on
// StoredBottleResponse
func ValidateStoredBottleResponse(body *StoredBottleResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Winery == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("winery", "body"))
	}
	if body.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("vintage", "body"))
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 100, false))
		}
	}
	if body.Winery != nil {
		if err2 := ValidateWineryResponse(body.Winery); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Vintage != nil {
		if *body.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 1900, true))
		}
	}
	if body.Vintage != nil {
		if *body.Vintage > 2020 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.vintage", *body.Vintage, 2020, false))
		}
	}
	for _, e := range body.Composition {
		if e != nil {
			if err2 := ValidateComponentResponse(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 2000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 2000, false))
		}
	}
	if body.Rating != nil {
		if *body.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 1, true))
		}
	}
	if body.Rating != nil {
		if *body.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.rating", *body.Rating, 5, false))
		}
	}
	return
}

// ValidateWineryResponse runs the validations defined on WineryResponse
func ValidateWineryResponse(body *WineryResponse) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Region == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("region", "body"))
	}
	if body.Country == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country", "body"))
	}
	if body.Region != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.region", *body.Region, "[a-zA-Z '\\.]+"))
	}
	if body.Country != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.country", *body.Country, "[a-zA-Z '\\.]+"))
	}
	if body.URL != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.url", *body.URL, "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$"))
	}
	return
}

// ValidateComponentResponse runs the validations defined on ComponentResponse
func ValidateComponentResponse(body *ComponentResponse) (err error) {
	if body.Varietal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("varietal", "body"))
	}
	if body.Varietal != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.varietal", *body.Varietal, "[A-Za-z' ]+"))
	}
	if body.Varietal != nil {
		if utf8.RuneCountInString(*body.Varietal) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.varietal", *body.Varietal, utf8.RuneCountInString(*body.Varietal), 100, false))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 1, true))
		}
	}
	if body.Percentage != nil {
		if *body.Percentage > 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.percentage", *body.Percentage, 100, false))
		}
	}
	return
}
