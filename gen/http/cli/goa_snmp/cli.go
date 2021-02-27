// Code generated by goa v3.2.6, DO NOT EDIT.
//
// goa-snmp HTTP client CLI support package
//
// Command:
// $ goa gen github.com/rshade/goasnmp/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	goasnmpc "github.com/rshade/goasnmp/gen/http/goa_snmp/client"
	goahttp "goa.design/goa/v3/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `goa-snmp (list|add)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` goa-snmp list` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (endpoint.Endpoint, interface{}, error) {
	var (
		goaSnmpFlags = flag.NewFlagSet("goa-snmp", flag.ContinueOnError)

		goaSnmpListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		goaSnmpAddFlags        = flag.NewFlagSet("add", flag.ExitOnError)
		goaSnmpAddBodyFlag     = goaSnmpAddFlags.String("body", "REQUIRED", "")
		goaSnmpAddHostnameFlag = goaSnmpAddFlags.String("hostname", "REQUIRED", "Hostname or Ip of Device")
	)
	goaSnmpFlags.Usage = goaSnmpUsage
	goaSnmpListFlags.Usage = goaSnmpListUsage
	goaSnmpAddFlags.Usage = goaSnmpAddUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "goa-snmp":
			svcf = goaSnmpFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "goa-snmp":
			switch epn {
			case "list":
				epf = goaSnmpListFlags

			case "add":
				epf = goaSnmpAddFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint endpoint.Endpoint
		err      error
	)
	{
		switch svcn {
		case "goa-snmp":
			c := goasnmpc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "add":
				endpoint = c.Add()
				data, err = goasnmpc.BuildAddPayload(*goaSnmpAddBodyFlag, *goaSnmpAddHostnameFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// goa-snmpUsage displays the usage of the goa-snmp command and its subcommands.
func goaSnmpUsage() {
	fmt.Fprintf(os.Stderr, `The goa-snmp service tracks hosts and walks snmp trees
Usage:
    %s [globalflags] goa-snmp COMMAND [flags]

COMMAND:
    list: List implements list.
    add: Add implements add.

Additional help:
    %s goa-snmp COMMAND --help
`, os.Args[0], os.Args[0])
}
func goaSnmpListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] goa-snmp list

List implements list.

Example:
    `+os.Args[0]+` goa-snmp list
`, os.Args[0])
}

func goaSnmpAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] goa-snmp add -body JSON -hostname STRING

Add implements add.
    -body JSON: 
    -hostname STRING: Hostname or Ip of Device

Example:
    `+os.Args[0]+` goa-snmp add --body '{
      "OnDemand": false,
      "Public": true
   }' --hostname "Voluptatum impedit saepe vitae qui."
`, os.Args[0])
}