// Code generated by goa v3.14.1, DO NOT EDIT.
//
// resume HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/multipart/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	resumec "goa.design/examples/multipart/gen/http/resume/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `resume (list|add)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` resume list` + "\n" +
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
	resumeAddEncoderFn resumec.ResumeAddEncoderFunc,
) (goa.Endpoint, any, error) {
	var (
		resumeFlags = flag.NewFlagSet("resume", flag.ContinueOnError)

		resumeListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		resumeAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		resumeAddBodyFlag = resumeAddFlags.String("body", "REQUIRED", "")
	)
	resumeFlags.Usage = resumeUsage
	resumeListFlags.Usage = resumeListUsage
	resumeAddFlags.Usage = resumeAddUsage

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
		case "resume":
			svcf = resumeFlags
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
		case "resume":
			switch epn {
			case "list":
				epf = resumeListFlags

			case "add":
				epf = resumeAddFlags

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
		case "resume":
			c := resumec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "add":
				endpoint = c.Add(resumeAddEncoderFn)
				data, err = resumec.BuildAddPayload(*resumeAddBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// resumeUsage displays the usage of the resume command and its subcommands.
func resumeUsage() {
	fmt.Fprintf(os.Stderr, `The storage service makes it possible to add resumes using multipart.
Usage:
    %[1]s [globalflags] resume COMMAND [flags]

COMMAND:
    list: List all stored resumes
    add: Add n number of resumes and return their IDs. This is a multipart request and each part has field name 'resume' and contains the encoded resume to be added.

Additional help:
    %[1]s resume COMMAND --help
`, os.Args[0])
}
func resumeListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] resume list

List all stored resumes

Example:
    %[1]s resume list
`, os.Args[0])
}

func resumeAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] resume add -body JSON

Add n number of resumes and return their IDs. This is a multipart request and each part has field name 'resume' and contains the encoded resume to be added.
    -body JSON: 

Example:
    %[1]s resume add --body '[
      {
         "education": [
            {
               "institution": "Nihil voluptate dolorem ut ipsam fuga.",
               "major": "Ex sequi."
            },
            {
               "institution": "Nihil voluptate dolorem ut ipsam fuga.",
               "major": "Ex sequi."
            }
         ],
         "experience": [
            {
               "company": "Saepe similique libero sunt.",
               "duration": 4538661945567210561,
               "role": "Cum maiores quo at ducimus sit."
            },
            {
               "company": "Saepe similique libero sunt.",
               "duration": 4538661945567210561,
               "role": "Cum maiores quo at ducimus sit."
            },
            {
               "company": "Saepe similique libero sunt.",
               "duration": 4538661945567210561,
               "role": "Cum maiores quo at ducimus sit."
            },
            {
               "company": "Saepe similique libero sunt.",
               "duration": 4538661945567210561,
               "role": "Cum maiores quo at ducimus sit."
            }
         ],
         "name": "In qui dolorum adipisci itaque."
      },
      {
         "education": [
            {
               "institution": "Nihil voluptate dolorem ut ipsam fuga.",
               "major": "Ex sequi."
            },
            {
               "institution": "Nihil voluptate dolorem ut ipsam fuga.",
               "major": "Ex sequi."
            }
         ],
         "experience": [
            {
               "company": "Saepe similique libero sunt.",
               "duration": 4538661945567210561,
               "role": "Cum maiores quo at ducimus sit."
            },
            {
               "company": "Saepe similique libero sunt.",
               "duration": 4538661945567210561,
               "role": "Cum maiores quo at ducimus sit."
            },
            {
               "company": "Saepe similique libero sunt.",
               "duration": 4538661945567210561,
               "role": "Cum maiores quo at ducimus sit."
            },
            {
               "company": "Saepe similique libero sunt.",
               "duration": 4538661945567210561,
               "role": "Cum maiores quo at ducimus sit."
            }
         ],
         "name": "In qui dolorum adipisci itaque."
      }
   ]'
`, os.Args[0])
}
