// Code generated by goa v3.11.2, DO NOT EDIT.
//
// chatter gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package cli

import (
	"flag"
	"fmt"
	"os"

	chatterc "goa.design/examples/streaming/gen/grpc/chatter/client"
	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `chatter (login|echoer|listener|summary|subscribe|history)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` chatter login --user "username" --password "password"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		chatterFlags = flag.NewFlagSet("chatter", flag.ContinueOnError)

		chatterLoginFlags        = flag.NewFlagSet("login", flag.ExitOnError)
		chatterLoginUserFlag     = chatterLoginFlags.String("user", "REQUIRED", "")
		chatterLoginPasswordFlag = chatterLoginFlags.String("password", "REQUIRED", "")

		chatterEchoerFlags     = flag.NewFlagSet("echoer", flag.ExitOnError)
		chatterEchoerTokenFlag = chatterEchoerFlags.String("token", "REQUIRED", "")

		chatterListenerFlags     = flag.NewFlagSet("listener", flag.ExitOnError)
		chatterListenerTokenFlag = chatterListenerFlags.String("token", "REQUIRED", "")

		chatterSummaryFlags     = flag.NewFlagSet("summary", flag.ExitOnError)
		chatterSummaryTokenFlag = chatterSummaryFlags.String("token", "REQUIRED", "")

		chatterSubscribeFlags     = flag.NewFlagSet("subscribe", flag.ExitOnError)
		chatterSubscribeTokenFlag = chatterSubscribeFlags.String("token", "REQUIRED", "")

		chatterHistoryFlags     = flag.NewFlagSet("history", flag.ExitOnError)
		chatterHistoryViewFlag  = chatterHistoryFlags.String("view", "", "")
		chatterHistoryTokenFlag = chatterHistoryFlags.String("token", "REQUIRED", "")
	)
	chatterFlags.Usage = chatterUsage
	chatterLoginFlags.Usage = chatterLoginUsage
	chatterEchoerFlags.Usage = chatterEchoerUsage
	chatterListenerFlags.Usage = chatterListenerUsage
	chatterSummaryFlags.Usage = chatterSummaryUsage
	chatterSubscribeFlags.Usage = chatterSubscribeUsage
	chatterHistoryFlags.Usage = chatterHistoryUsage

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
		case "chatter":
			svcf = chatterFlags
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
		case "chatter":
			switch epn {
			case "login":
				epf = chatterLoginFlags

			case "echoer":
				epf = chatterEchoerFlags

			case "listener":
				epf = chatterListenerFlags

			case "summary":
				epf = chatterSummaryFlags

			case "subscribe":
				epf = chatterSubscribeFlags

			case "history":
				epf = chatterHistoryFlags

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
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "chatter":
			c := chatterc.NewClient(cc, opts...)
			switch epn {
			case "login":
				endpoint = c.Login()
				data, err = chatterc.BuildLoginPayload(*chatterLoginUserFlag, *chatterLoginPasswordFlag)
			case "echoer":
				endpoint = c.Echoer()
				data, err = chatterc.BuildEchoerPayload(*chatterEchoerTokenFlag)
			case "listener":
				endpoint = c.Listener()
				data, err = chatterc.BuildListenerPayload(*chatterListenerTokenFlag)
			case "summary":
				endpoint = c.Summary()
				data, err = chatterc.BuildSummaryPayload(*chatterSummaryTokenFlag)
			case "subscribe":
				endpoint = c.Subscribe()
				data, err = chatterc.BuildSubscribePayload(*chatterSubscribeTokenFlag)
			case "history":
				endpoint = c.History()
				data, err = chatterc.BuildHistoryPayload(*chatterHistoryViewFlag, *chatterHistoryTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// chatterUsage displays the usage of the chatter command and its subcommands.
func chatterUsage() {
	fmt.Fprintf(os.Stderr, `The chatter service implements a simple client and server chat.
Usage:
    %[1]s [globalflags] chatter COMMAND [flags]

COMMAND:
    login: Creates a valid JWT token for auth to chat.
    echoer: Echoes the message sent by the client.
    listener: Listens to the messages sent by the client.
    summary: Summarizes the chat messages sent by the client.
    subscribe: Subscribe to events sent when new chat messages are added.
    history: Returns the chat messages sent to the server.

Additional help:
    %[1]s chatter COMMAND --help
`, os.Args[0])
}
func chatterLoginUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chatter login -user STRING -password STRING

Creates a valid JWT token for auth to chat.
    -user STRING: 
    -password STRING: 

Example:
    %[1]s chatter login --user "username" --password "password"
`, os.Args[0])
}

func chatterEchoerUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chatter echoer -token STRING

Echoes the message sent by the client.
    -token STRING: 

Example:
    %[1]s chatter echoer --token "Fugit et id cum cum."
`, os.Args[0])
}

func chatterListenerUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chatter listener -token STRING

Listens to the messages sent by the client.
    -token STRING: 

Example:
    %[1]s chatter listener --token "Perspiciatis sed tempore."
`, os.Args[0])
}

func chatterSummaryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chatter summary -token STRING

Summarizes the chat messages sent by the client.
    -token STRING: 

Example:
    %[1]s chatter summary --token "Et soluta aut tempora autem sint."
`, os.Args[0])
}

func chatterSubscribeUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chatter subscribe -token STRING

Subscribe to events sent when new chat messages are added.
    -token STRING: 

Example:
    %[1]s chatter subscribe --token "Ut dolor repellat blanditiis quia qui aspernatur."
`, os.Args[0])
}

func chatterHistoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] chatter history -view STRING -token STRING

Returns the chat messages sent to the server.
    -view STRING: 
    -token STRING: 

Example:
    %[1]s chatter history --view "Odit dolor dolorum reprehenderit dolorum voluptas." --token "Repudiandae ipsa aut quaerat."
`, os.Args[0])
}
