package goasnmpapi

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	goasnmp "github.com/rshade/goasnmp/gen/goasnmp"
)

// goasnmp service example implementation.
// The example methods log the requests and return zero values.
type goasnmpsrvc struct {
	logger log.Logger
}

// NewGoasnmp returns the goasnmp service implementation.
func NewGoasnmp(logger log.Logger) goasnmp.Service {
	return &goasnmpsrvc{logger}
}

// List implements list.
func (s *goasnmpsrvc) List(ctx context.Context) (res goasnmp.HostCollection, err error) {
	s.logger.Log("info", fmt.Sprintf("goasnmp.list"))
	return
}

// Add implements add.
func (s *goasnmpsrvc) Add(ctx context.Context, p *goasnmp.AddPayload) (res *goasnmp.Host, err error) {
	res = &goasnmp.Host{}
	s.logger.Log("info", fmt.Sprintf("goasnmp.add"))
	return
}
