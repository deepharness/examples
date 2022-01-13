// Code generated by goa v3.5.4, DO NOT EDIT.
//
// sommelier service
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o
// $(GOPATH)/src/goa.design/examples/cellar

package sommelier

import (
	"context"

	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
)

// The sommelier service retrieves bottles given a set of criteria.
type Service interface {
	// Pick implements pick.
	Pick(context.Context, *Criteria) (res StoredBottleCollection, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "sommelier"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"pick"}

// Criteria is the payload type of the sommelier service pick method.
type Criteria struct {
	// Name of bottle to pick
	Name *string
	// Varietals in preference order
	Varietal []string
	// Winery of bottle to pick
	Winery *string
}

// StoredBottleCollection is the result type of the sommelier service pick
// method.
type StoredBottleCollection []*StoredBottle

// A StoredBottle describes a bottle retrieved by the storage service.
type StoredBottle struct {
	// ID is the unique id of the bottle.
	ID string
	// Name of bottle
	Name string
	// Winery that produces wine
	Winery *Winery
	// Vintage of bottle
	Vintage uint32
	// Composition is the list of grape varietals and associated percentage.
	Composition []*Component
	// Description of bottle
	Description *string
	// Rating of bottle from 1 (worst) to 5 (best)
	Rating *uint32
}

type Winery struct {
	// Name of winery
	Name string
	// Region of winery
	Region string
	// Country of winery
	Country string
	// Winery website URL
	URL *string
}

type Component struct {
	// Grape varietal
	Varietal string
	// Percentage of varietal in wine
	Percentage *uint32
}

// Missing criteria
type NoCriteria string

// No bottle matched given criteria
type NoMatch string

// Error returns an error description.
func (e NoCriteria) Error() string {
	return "Missing criteria"
}

// ErrorName returns "no_criteria".
func (e NoCriteria) ErrorName() string {
	return "no_criteria"
}

// Error returns an error description.
func (e NoMatch) Error() string {
	return "No bottle matched given criteria"
}

// ErrorName returns "no_match".
func (e NoMatch) ErrorName() string {
	return "no_match"
}

// NewStoredBottleCollection initializes result type StoredBottleCollection
// from viewed result type StoredBottleCollection.
func NewStoredBottleCollection(vres sommelierviews.StoredBottleCollection) StoredBottleCollection {
	var res StoredBottleCollection
	switch vres.View {
	case "default", "":
		res = newStoredBottleCollection(vres.Projected)
	case "tiny":
		res = newStoredBottleCollectionTiny(vres.Projected)
	}
	return res
}

// NewViewedStoredBottleCollection initializes viewed result type
// StoredBottleCollection from result type StoredBottleCollection using the
// given view.
func NewViewedStoredBottleCollection(res StoredBottleCollection, view string) sommelierviews.StoredBottleCollection {
	var vres sommelierviews.StoredBottleCollection
	switch view {
	case "default", "":
		p := newStoredBottleCollectionView(res)
		vres = sommelierviews.StoredBottleCollection{Projected: p, View: "default"}
	case "tiny":
		p := newStoredBottleCollectionViewTiny(res)
		vres = sommelierviews.StoredBottleCollection{Projected: p, View: "tiny"}
	}
	return vres
}

// newStoredBottleCollection converts projected type StoredBottleCollection to
// service type StoredBottleCollection.
func newStoredBottleCollection(vres sommelierviews.StoredBottleCollectionView) StoredBottleCollection {
	res := make(StoredBottleCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredBottle(n)
	}
	return res
}

// newStoredBottleCollectionTiny converts projected type StoredBottleCollection
// to service type StoredBottleCollection.
func newStoredBottleCollectionTiny(vres sommelierviews.StoredBottleCollectionView) StoredBottleCollection {
	res := make(StoredBottleCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredBottleTiny(n)
	}
	return res
}

// newStoredBottleCollectionView projects result type StoredBottleCollection to
// projected type StoredBottleCollectionView using the "default" view.
func newStoredBottleCollectionView(res StoredBottleCollection) sommelierviews.StoredBottleCollectionView {
	vres := make(sommelierviews.StoredBottleCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredBottleView(n)
	}
	return vres
}

// newStoredBottleCollectionViewTiny projects result type
// StoredBottleCollection to projected type StoredBottleCollectionView using
// the "tiny" view.
func newStoredBottleCollectionViewTiny(res StoredBottleCollection) sommelierviews.StoredBottleCollectionView {
	vres := make(sommelierviews.StoredBottleCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredBottleViewTiny(n)
	}
	return vres
}

// newStoredBottle converts projected type StoredBottle to service type
// StoredBottle.
func newStoredBottle(vres *sommelierviews.StoredBottleView) *StoredBottle {
	res := &StoredBottle{
		Description: vres.Description,
		Rating:      vres.Rating,
	}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Vintage != nil {
		res.Vintage = *vres.Vintage
	}
	if vres.Composition != nil {
		res.Composition = make([]*Component, len(vres.Composition))
		for i, val := range vres.Composition {
			res.Composition[i] = transformSommelierviewsComponentViewToComponent(val)
		}
	}
	if vres.Winery != nil {
		res.Winery = newWineryTiny(vres.Winery)
	}
	return res
}

// newStoredBottleTiny converts projected type StoredBottle to service type
// StoredBottle.
func newStoredBottleTiny(vres *sommelierviews.StoredBottleView) *StoredBottle {
	res := &StoredBottle{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Winery != nil {
		res.Winery = newWineryTiny(vres.Winery)
	}
	return res
}

// newStoredBottleView projects result type StoredBottle to projected type
// StoredBottleView using the "default" view.
func newStoredBottleView(res *StoredBottle) *sommelierviews.StoredBottleView {
	vres := &sommelierviews.StoredBottleView{
		ID:          &res.ID,
		Name:        &res.Name,
		Vintage:     &res.Vintage,
		Description: res.Description,
		Rating:      res.Rating,
	}
	if res.Composition != nil {
		vres.Composition = make([]*sommelierviews.ComponentView, len(res.Composition))
		for i, val := range res.Composition {
			vres.Composition[i] = transformComponentToSommelierviewsComponentView(val)
		}
	}
	if res.Winery != nil {
		vres.Winery = newWineryViewTiny(res.Winery)
	}
	return vres
}

// newStoredBottleViewTiny projects result type StoredBottle to projected type
// StoredBottleView using the "tiny" view.
func newStoredBottleViewTiny(res *StoredBottle) *sommelierviews.StoredBottleView {
	vres := &sommelierviews.StoredBottleView{
		ID:   &res.ID,
		Name: &res.Name,
	}
	if res.Winery != nil {
		vres.Winery = newWineryViewTiny(res.Winery)
	}
	return vres
}

// newWinery converts projected type Winery to service type Winery.
func newWinery(vres *sommelierviews.WineryView) *Winery {
	res := &Winery{
		URL: vres.URL,
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Region != nil {
		res.Region = *vres.Region
	}
	if vres.Country != nil {
		res.Country = *vres.Country
	}
	return res
}

// newWineryTiny converts projected type Winery to service type Winery.
func newWineryTiny(vres *sommelierviews.WineryView) *Winery {
	res := &Winery{}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	return res
}

// newWineryView projects result type Winery to projected type WineryView using
// the "default" view.
func newWineryView(res *Winery) *sommelierviews.WineryView {
	vres := &sommelierviews.WineryView{
		Name:    &res.Name,
		Region:  &res.Region,
		Country: &res.Country,
		URL:     res.URL,
	}
	return vres
}

// newWineryViewTiny projects result type Winery to projected type WineryView
// using the "tiny" view.
func newWineryViewTiny(res *Winery) *sommelierviews.WineryView {
	vres := &sommelierviews.WineryView{
		Name: &res.Name,
	}
	return vres
}

// transformSommelierviewsComponentViewToComponent builds a value of type
// *Component from a value of type *sommelierviews.ComponentView.
func transformSommelierviewsComponentViewToComponent(v *sommelierviews.ComponentView) *Component {
	if v == nil {
		return nil
	}
	res := &Component{
		Varietal:   *v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}

// transformComponentToSommelierviewsComponentView builds a value of type
// *sommelierviews.ComponentView from a value of type *Component.
func transformComponentToSommelierviewsComponentView(v *Component) *sommelierviews.ComponentView {
	if v == nil {
		return nil
	}
	res := &sommelierviews.ComponentView{
		Varietal:   &v.Varietal,
		Percentage: v.Percentage,
	}

	return res
}
