package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
	media_type "github.com/rshade/goasnmp/design/media_type"
)

var _ = API("goasnmp", func() {
	Title("Goa SNMP Service")
	Description("Service for walking snmp trees and converting to http")
})

var _ = Service("goasnmp", func() {
	Description("The goa-snmp service tracks hosts and walks snmp trees")

	Method("list", func() {
		Result(CollectionOf(media_type.Host))
		HTTP(func() {
			GET("/hosts")
		})
	})

	Method("add", func() {
		Payload(func() {
			Field(1, "Hostname", String, "Hostname or Ip of Device")
			Field(2, "Public", Boolean, "Whether or not to walk public tree")
			Field(3, "OnDemand", Boolean, "Whether or not Ondemand polling is supported")
			Required("Hostname", "Public")
		})

		Result(media_type.Host)

		HTTP(func() {
			POST("/hosts/{Hostname}")
			Body(func() {
				Attribute("Public", Boolean, "Whether or not to walk public tree")
				Attribute("OnDemand", Boolean, "Whether or not Ondemand polling is supported")
			})
			Response(StatusOK)
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
