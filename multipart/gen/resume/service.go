// Code generated by goa v3.9.1, DO NOT EDIT.
//
// resume service
//
// Command:
// $ goa gen goa.design/examples/multipart/design -o multipart

package resume

import (
	"context"

	resumeviews "goa.design/examples/multipart/gen/resume/views"
)

// The storage service makes it possible to add resumes using multipart.
type Service interface {
	// List all stored resumes
	List(context.Context) (res StoredResumeCollection, err error)
	// Add n number of resumes and return their IDs. This is a multipart request
	// and each part has field name 'resume' and contains the encoded resume to be
	// added.
	Add(context.Context, []*Resume) (res []int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "resume"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"list", "add"}

type Education struct {
	// Name of the institution
	Institution string
	// Major name
	Major string
}

type Experience struct {
	// Name of the company
	Company string
	// Name of the role in the company
	Role string
	// Duration (in years) in the company
	Duration int
}

type Resume struct {
	// Name in the resume
	Name string
	// Experience section in the resume
	Experience []*Experience
	// Education section in the resume
	Education []*Education
}

type StoredResume struct {
	// ID of the resume
	ID int
	// Time when resume was created
	CreatedAt string
	// Name in the resume
	Name string
	// Experience section in the resume
	Experience []*Experience
	// Education section in the resume
	Education []*Education
}

// StoredResumeCollection is the result type of the resume service list method.
type StoredResumeCollection []*StoredResume

// NewStoredResumeCollection initializes result type StoredResumeCollection
// from viewed result type StoredResumeCollection.
func NewStoredResumeCollection(vres resumeviews.StoredResumeCollection) StoredResumeCollection {
	return newStoredResumeCollection(vres.Projected)
}

// NewViewedStoredResumeCollection initializes viewed result type
// StoredResumeCollection from result type StoredResumeCollection using the
// given view.
func NewViewedStoredResumeCollection(res StoredResumeCollection, view string) resumeviews.StoredResumeCollection {
	p := newStoredResumeCollectionView(res)
	return resumeviews.StoredResumeCollection{Projected: p, View: "default"}
}

// newStoredResumeCollection converts projected type StoredResumeCollection to
// service type StoredResumeCollection.
func newStoredResumeCollection(vres resumeviews.StoredResumeCollectionView) StoredResumeCollection {
	res := make(StoredResumeCollection, len(vres))
	for i, n := range vres {
		res[i] = newStoredResume(n)
	}
	return res
}

// newStoredResumeCollectionView projects result type StoredResumeCollection to
// projected type StoredResumeCollectionView using the "default" view.
func newStoredResumeCollectionView(res StoredResumeCollection) resumeviews.StoredResumeCollectionView {
	vres := make(resumeviews.StoredResumeCollectionView, len(res))
	for i, n := range res {
		vres[i] = newStoredResumeView(n)
	}
	return vres
}

// newStoredResume converts projected type StoredResume to service type
// StoredResume.
func newStoredResume(vres *resumeviews.StoredResumeView) *StoredResume {
	res := &StoredResume{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.CreatedAt != nil {
		res.CreatedAt = *vres.CreatedAt
	}
	if vres.Experience != nil {
		res.Experience = make([]*Experience, len(vres.Experience))
		for i, val := range vres.Experience {
			res.Experience[i] = transformResumeviewsExperienceViewToExperience(val)
		}
	}
	if vres.Education != nil {
		res.Education = make([]*Education, len(vres.Education))
		for i, val := range vres.Education {
			res.Education[i] = transformResumeviewsEducationViewToEducation(val)
		}
	}
	return res
}

// newStoredResumeView projects result type StoredResume to projected type
// StoredResumeView using the "default" view.
func newStoredResumeView(res *StoredResume) *resumeviews.StoredResumeView {
	vres := &resumeviews.StoredResumeView{
		ID:        &res.ID,
		CreatedAt: &res.CreatedAt,
		Name:      &res.Name,
	}
	if res.Experience != nil {
		vres.Experience = make([]*resumeviews.ExperienceView, len(res.Experience))
		for i, val := range res.Experience {
			vres.Experience[i] = transformExperienceToResumeviewsExperienceView(val)
		}
	}
	if res.Education != nil {
		vres.Education = make([]*resumeviews.EducationView, len(res.Education))
		for i, val := range res.Education {
			vres.Education[i] = transformEducationToResumeviewsEducationView(val)
		}
	}
	return vres
}

// transformResumeviewsExperienceViewToExperience builds a value of type
// *Experience from a value of type *resumeviews.ExperienceView.
func transformResumeviewsExperienceViewToExperience(v *resumeviews.ExperienceView) *Experience {
	if v == nil {
		return nil
	}
	res := &Experience{
		Company:  *v.Company,
		Role:     *v.Role,
		Duration: *v.Duration,
	}

	return res
}

// transformResumeviewsEducationViewToEducation builds a value of type
// *Education from a value of type *resumeviews.EducationView.
func transformResumeviewsEducationViewToEducation(v *resumeviews.EducationView) *Education {
	if v == nil {
		return nil
	}
	res := &Education{
		Institution: *v.Institution,
		Major:       *v.Major,
	}

	return res
}

// transformExperienceToResumeviewsExperienceView builds a value of type
// *resumeviews.ExperienceView from a value of type *Experience.
func transformExperienceToResumeviewsExperienceView(v *Experience) *resumeviews.ExperienceView {
	res := &resumeviews.ExperienceView{
		Company:  &v.Company,
		Role:     &v.Role,
		Duration: &v.Duration,
	}

	return res
}

// transformEducationToResumeviewsEducationView builds a value of type
// *resumeviews.EducationView from a value of type *Education.
func transformEducationToResumeviewsEducationView(v *Education) *resumeviews.EducationView {
	res := &resumeviews.EducationView{
		Institution: &v.Institution,
		Major:       &v.Major,
	}

	return res
}
