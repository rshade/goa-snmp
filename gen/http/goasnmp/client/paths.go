// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the goasnmp service.
//
// Command:
// $ goa gen github.com/rshade/goasnmp/design

package client

import (
	"fmt"
)

// ListGoasnmpPath returns the URL path to the goasnmp service list HTTP endpoint.
func ListGoasnmpPath() string {
	return "/hosts"
}

// AddGoasnmpPath returns the URL path to the goasnmp service add HTTP endpoint.
func AddGoasnmpPath(hostname string) string {
	return fmt.Sprintf("/hosts/%v", hostname)
}
