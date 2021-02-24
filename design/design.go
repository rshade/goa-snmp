package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/goakit"
)

type InterfaceHost struct {
	Name string `json:"name"` // name or ip of device
	Public bool `json:"public"` // walk the public tree
	OnDemand bool `json:"onDemand"` //supports ondemand vs cached data
	Href string `json:"href"`
}

var _ = API("goa-snmp", func() {
	Title("Goa SNMP Service")
	Description("Service for walking snmp trees and converting to http")
    Server("goa-snmp", func() {
        Host("localhost", func() {
            URI("http://localhost:8000")
            URI("grpc://localhost:8080")
        })
    })
})

var _ = Service("goa-snmp", func() {
	Description("The goa-snmp service tracks hosts and walks snmp trees")

	Method("list", func(){
		Result(CollectionOf(InterfaceHost{}))
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

		Result(InterfaceHost{})

		HTTP(func() {
			POST("/hosts/{Hostname}")
			Body(func(){
				Attribute("Public", Boolean, "Whether or not to walk public tree")
				Attribute("OnDemand", Boolean, "Whether or not Ondemand polling is supported")
			})
			Response(StatusOK)
		})

		GRPC(func() {
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})