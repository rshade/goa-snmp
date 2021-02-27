package media_type

import (
	. "goa.design/goa/v3/dsl"
)

type InterfaceHost struct {
	Name     string `json:"name"`     // name or ip of device
	Public   bool   `json:"public"`   // walk the public tree
	OnDemand bool   `json:"onDemand"` //supports ondemand vs cached data
	Href     string `json:"href"`
}

var SnmpHost = ResultType("application/vnd.goasnmp.host", func() {
	TypeName("Host")
	Description("Host to be queried")
	CreateFrom(InterfaceHost{})
	Attributes(func() {
		Attribute("public", Boolean, "Whether or not to walk public tree")
		Attribute("on_demand", Boolean, "Whether or not Ondemand polling is supported")
	})
})
