package design

import (
	d "github.com/goadesign/goa/design"
	a "github.com/goadesign/goa/design/apidsl"
)

var _ = a.Resource("workitem_refresh", func() {
	a.Parent("workitem")
	a.Action("refresh", func() {
		a.Routing(a.GET("refresh"))
		a.Scheme("ws")
		a.Description("refresh work item")
		a.Response(d.SwitchingProtocols)
	})

})
