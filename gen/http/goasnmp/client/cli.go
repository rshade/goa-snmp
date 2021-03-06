// Code generated by goa v3.2.6, DO NOT EDIT.
//
// goasnmp HTTP client CLI support package
//
// Command:
// $ goa gen github.com/rshade/goasnmp/design

package client

import (
	"encoding/json"
	"fmt"

	goasnmp "github.com/rshade/goasnmp/gen/goasnmp"
)

// BuildAddPayload builds the payload for the goasnmp add endpoint from CLI
// flags.
func BuildAddPayload(goasnmpAddBody string, goasnmpAddHostname string) (*goasnmp.AddPayload, error) {
	var err error
	var body struct {
		// Whether or not to walk public tree
		Public *bool `form:"Public" json:"Public" xml:"Public"`
		// Whether or not Ondemand polling is supported
		OnDemand *bool `form:"OnDemand" json:"OnDemand" xml:"OnDemand"`
	}
	{
		err = json.Unmarshal([]byte(goasnmpAddBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"OnDemand\": false,\n      \"Public\": true\n   }'")
		}
	}
	var hostname string
	{
		hostname = goasnmpAddHostname
	}
	v := &goasnmp.AddPayload{
		OnDemand: body.OnDemand,
	}
	if body.Public != nil {
		v.Public = *body.Public
	}
	v.Hostname = hostname

	return v, nil
}
