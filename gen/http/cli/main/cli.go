// Code generated by goa v3.20.1, DO NOT EDIT.
//
// main HTTP client CLI support package
//
// Command:
// $ goa gen github.com/meier-christoph/todo-backend-golang-goa/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	todosc "github.com/meier-christoph/todo-backend-golang-goa/gen/http/todos/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `todos (list|create|read|update|delete|delete-all)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` todos list` + "\n" +
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
) (goa.Endpoint, any, error) {
	var (
		todosFlags = flag.NewFlagSet("todos", flag.ContinueOnError)

		todosListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		todosCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		todosCreateBodyFlag = todosCreateFlags.String("body", "REQUIRED", "")

		todosReadFlags  = flag.NewFlagSet("read", flag.ExitOnError)
		todosReadIDFlag = todosReadFlags.String("id", "REQUIRED", "")

		todosUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		todosUpdateBodyFlag = todosUpdateFlags.String("body", "REQUIRED", "")
		todosUpdateIDFlag   = todosUpdateFlags.String("id", "REQUIRED", "")

		todosDeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		todosDeleteIDFlag = todosDeleteFlags.String("id", "REQUIRED", "")

		todosDeleteAllFlags = flag.NewFlagSet("delete-all", flag.ExitOnError)
	)
	todosFlags.Usage = todosUsage
	todosListFlags.Usage = todosListUsage
	todosCreateFlags.Usage = todosCreateUsage
	todosReadFlags.Usage = todosReadUsage
	todosUpdateFlags.Usage = todosUpdateUsage
	todosDeleteFlags.Usage = todosDeleteUsage
	todosDeleteAllFlags.Usage = todosDeleteAllUsage

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
		case "todos":
			svcf = todosFlags
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
		case "todos":
			switch epn {
			case "list":
				epf = todosListFlags

			case "create":
				epf = todosCreateFlags

			case "read":
				epf = todosReadFlags

			case "update":
				epf = todosUpdateFlags

			case "delete":
				epf = todosDeleteFlags

			case "delete-all":
				epf = todosDeleteAllFlags

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
		data     any
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "todos":
			c := todosc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
			case "create":
				endpoint = c.Create()
				data, err = todosc.BuildCreatePayload(*todosCreateBodyFlag)
			case "read":
				endpoint = c.Read()
				data, err = todosc.BuildReadPayload(*todosReadIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = todosc.BuildUpdatePayload(*todosUpdateBodyFlag, *todosUpdateIDFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = todosc.BuildDeletePayload(*todosDeleteIDFlag)
			case "delete-all":
				endpoint = c.DeleteAll()
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// todosUsage displays the usage of the todos command and its subcommands.
func todosUsage() {
	fmt.Fprintf(os.Stderr, `Service is the todos service interface.
Usage:
    %[1]s [globalflags] todos COMMAND [flags]

COMMAND:
    list: List implements list.
    create: Create implements create.
    read: Read implements read.
    update: Update implements update.
    delete: Delete implements delete.
    delete-all: DeleteAll implements deleteAll.

Additional help:
    %[1]s todos COMMAND --help
`, os.Args[0])
}
func todosListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos list

List implements list.

Example:
    %[1]s todos list
`, os.Args[0])
}

func todosCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos create -body JSON

Create implements create.
    -body JSON: 

Example:
    %[1]s todos create --body '{
      "completed": false,
      "order": 3935193967633306166,
      "title": "Voluptatibus adipisci eos vel."
   }'
`, os.Args[0])
}

func todosReadUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos read -id STRING

Read implements read.
    -id STRING: 

Example:
    %[1]s todos read --id "Ab iure dolore voluptatum est."
`, os.Args[0])
}

func todosUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos update -body JSON -id STRING

Update implements update.
    -body JSON: 
    -id STRING: 

Example:
    %[1]s todos update --body '{
      "completed": true,
      "order": 4610252795505360349,
      "title": "Molestiae molestias rem exercitationem vel officia consequatur."
   }' --id "Voluptas qui qui aut rerum eos quia."
`, os.Args[0])
}

func todosDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos delete -id STRING

Delete implements delete.
    -id STRING: 

Example:
    %[1]s todos delete --id "Sit voluptatem officia est neque ex."
`, os.Args[0])
}

func todosDeleteAllUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] todos delete-all

DeleteAll implements deleteAll.

Example:
    %[1]s todos delete-all
`, os.Args[0])
}
