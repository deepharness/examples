package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Secured API", func() {
	Description("This artifical API shows 4 different ways to secure API endpoints: using " +
		"basic auth, shared secret header, JWT or OAuth2")
})

var SuccessMedia = MediaType("application/vnd.goa.examples.security.success", func() {
	Description("The common media type to all request responses for this example")
	TypeName("Success")
	Attributes(func() {
		Attribute("ok", Boolean, "Always true")
	})
	View("default", func() {
		Attribute("ok")
	})
})
