package goasnmpapi

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	goasnmp "github.com/rshade/goasnmp/gen/goa_snmp"
)

// goa-snmp service example implementation.
// The example methods log the requests and return zero values.
type goaSnmpsrvc struct {
	logger log.Logger
}

// NewGoaSnmp returns the goa-snmp service implementation.
func NewGoaSnmp(logger log.Logger) goasnmp.Service {
	return &goaSnmpsrvc{logger}
}

// List implements list.
func (s *goaSnmpsrvc) List(ctx context.Context) (res goasnmp.HostCollection, err error) {
	s.logger.Log("info", fmt.Sprintf("goaSnmp.list"))
	return
}

// Add implements add.
func (s *goaSnmpsrvc) Add(ctx context.Context, p *goasnmp.AddPayload) (res *goasnmp.Host, err error) {
	res = &goasnmp.Host{}
	s.logger.Log("info", fmt.Sprintf("goaSnmp.add"))
	return
}
