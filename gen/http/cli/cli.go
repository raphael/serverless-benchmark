// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// recorder HTTP client CLI support package
//
// Command:
// $ goa gen github.com/raphael/recorder/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	recordersvcc "github.com/raphael/recorder/gen/http/recorder/client"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `recorder (record-data|list)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` recorder record-data --body '{
      "name": "duration",
      "service": "lambda",
      "value": 0.06173619203715241
   }'` + "\n" +
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
) (goa.Endpoint, interface{}, error) {
	var (
		recorderFlags = flag.NewFlagSet("recorder", flag.ContinueOnError)

		recorderRecordDataFlags    = flag.NewFlagSet("record-data", flag.ExitOnError)
		recorderRecordDataBodyFlag = recorderRecordDataFlags.String("body", "REQUIRED", "")

		recorderListFlags       = flag.NewFlagSet("list", flag.ExitOnError)
		recorderListServiceFlag = recorderListFlags.String("service", "REQUIRED", "")
		recorderListNameFlag    = recorderListFlags.String("name", "REQUIRED", "")
	)
	recorderFlags.Usage = recorderUsage
	recorderRecordDataFlags.Usage = recorderRecordDataUsage
	recorderListFlags.Usage = recorderListUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if len(os.Args) < flag.NFlag()+3 {
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = os.Args[1+flag.NFlag()]
		switch svcn {
		case "recorder":
			svcf = recorderFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(os.Args[2+flag.NFlag():]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = os.Args[2+flag.NFlag()+svcf.NFlag()]
		switch svcn {
		case "recorder":
			switch epn {
			case "record-data":
				epf = recorderRecordDataFlags

			case "list":
				epf = recorderListFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if len(os.Args) > 2+flag.NFlag()+svcf.NFlag() {
		if err := epf.Parse(os.Args[3+flag.NFlag()+svcf.NFlag():]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "recorder":
			c := recordersvcc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "record-data":
				endpoint = c.RecordData()
				data, err = recordersvcc.BuildRecordDataPayload(*recorderRecordDataBodyFlag)
			case "list":
				endpoint = c.List()
				data, err = recordersvcc.BuildListPayload(*recorderListServiceFlag, *recorderListNameFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// recorderUsage displays the usage of the recorder command and its subcommands.
func recorderUsage() {
	fmt.Fprintf(os.Stderr, `Service is the recorder service interface.
Usage:
    %s [globalflags] recorder COMMAND [flags]

COMMAND:
    record-data: RecordData creates a new datapoint.
    list: List lists all recorded datapoints.

Additional help:
    %s recorder COMMAND --help
`, os.Args[0], os.Args[0])
}
func recorderRecordDataUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] recorder record-data -body JSON

RecordData creates a new datapoint.
    -body JSON: 

Example:
    `+os.Args[0]+` recorder record-data --body '{
      "name": "duration",
      "service": "lambda",
      "value": 0.06173619203715241
   }'
`, os.Args[0])
}

func recorderListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] recorder list -service STRING -name STRING

List lists all recorded datapoints.
    -service STRING: 
    -name STRING: 

Example:
    `+os.Args[0]+` recorder list --service "lambda" --name "duration"
`, os.Args[0])
}
