// Code generated by goa v3.16.1, DO NOT EDIT.
//
// cellar HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/cellar/design

package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	sommelierc "goa.design/examples/cellar/gen/http/sommelier/client"
	storagec "goa.design/examples/cellar/gen/http/storage/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `sommelier pick
storage (list|show|add|remove|rate|multi-add|multi-update)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` sommelier pick --body '{
      "name": "Blue\'s Cuvee",
      "varietal": [
         "pinot noir",
         "merlot",
         "cabernet franc"
      ],
      "winery": "longoria"
   }'` + "\n" +
		os.Args[0] + ` storage list` + "\n" +
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
	storageMultiAddEncoderFn storagec.StorageMultiAddEncoderFunc,
	storageMultiUpdateEncoderFn storagec.StorageMultiUpdateEncoderFunc,
) (goa.Endpoint, any, error) {
	var (
		sommelierFlags = flag.NewFlagSet("sommelier", flag.ContinueOnError)

		sommelierPickFlags    = flag.NewFlagSet("pick", flag.ExitOnError)
		sommelierPickBodyFlag = sommelierPickFlags.String("body", "REQUIRED", "")

		storageFlags = flag.NewFlagSet("storage", flag.ContinueOnError)

		storageListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		storageShowFlags    = flag.NewFlagSet("show", flag.ExitOnError)
		storageShowIDFlag   = storageShowFlags.String("id", "REQUIRED", "ID of bottle to show")
		storageShowViewFlag = storageShowFlags.String("view", "", "")

		storageAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		storageAddBodyFlag = storageAddFlags.String("body", "REQUIRED", "")

		storageRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		storageRemoveIDFlag = storageRemoveFlags.String("id", "REQUIRED", "ID of bottle to remove")

		storageRateFlags = flag.NewFlagSet("rate", flag.ExitOnError)
		storageRatePFlag = storageRateFlags.String("p", "REQUIRED", "map[uint32][]string is the payload type of the storage service rate method.")

		storageMultiAddFlags    = flag.NewFlagSet("multi-add", flag.ExitOnError)
		storageMultiAddBodyFlag = storageMultiAddFlags.String("body", "REQUIRED", "")

		storageMultiUpdateFlags    = flag.NewFlagSet("multi-update", flag.ExitOnError)
		storageMultiUpdateBodyFlag = storageMultiUpdateFlags.String("body", "REQUIRED", "")
		storageMultiUpdateIdsFlag  = storageMultiUpdateFlags.String("ids", "REQUIRED", "")
	)
	sommelierFlags.Usage = sommelierUsage
	sommelierPickFlags.Usage = sommelierPickUsage

	storageFlags.Usage = storageUsage
	storageListFlags.Usage = storageListUsage
	storageShowFlags.Usage = storageShowUsage
	storageAddFlags.Usage = storageAddUsage
	storageRemoveFlags.Usage = storageRemoveUsage
	storageRateFlags.Usage = storageRateUsage
	storageMultiAddFlags.Usage = storageMultiAddUsage
	storageMultiUpdateFlags.Usage = storageMultiUpdateUsage

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
		case "sommelier":
			svcf = sommelierFlags
		case "storage":
			svcf = storageFlags
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
		case "sommelier":
			switch epn {
			case "pick":
				epf = sommelierPickFlags

			}

		case "storage":
			switch epn {
			case "list":
				epf = storageListFlags

			case "show":
				epf = storageShowFlags

			case "add":
				epf = storageAddFlags

			case "remove":
				epf = storageRemoveFlags

			case "rate":
				epf = storageRateFlags

			case "multi-add":
				epf = storageMultiAddFlags

			case "multi-update":
				epf = storageMultiUpdateFlags

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
		case "sommelier":
			c := sommelierc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "pick":
				endpoint = c.Pick()
				data, err = sommelierc.BuildPickPayload(*sommelierPickBodyFlag)
			}
		case "storage":
			c := storagec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = storagec.BuildShowPayload(*storageShowIDFlag, *storageShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = storagec.BuildAddPayload(*storageAddBodyFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = storagec.BuildRemovePayload(*storageRemoveIDFlag)
			case "rate":
				endpoint = c.Rate()
				var err error
				var val map[uint32][]string
				err = json.Unmarshal([]byte(*storageRatePFlag), &val)
				data = val
				if err != nil {
					return nil, nil, fmt.Errorf("invalid JSON for storageRatePFlag, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"\\u003cuint32 Value\\u003e\": [\n         \"Laboriosam consequatur delectus doloribus.\",\n         \"Est mollitia.\",\n         \"Voluptas ex enim.\",\n         \"Est explicabo eveniet dolore.\"\n      ]\n   }'")
				}
			case "multi-add":
				endpoint = c.MultiAdd(storageMultiAddEncoderFn)
				data, err = storagec.BuildMultiAddPayload(*storageMultiAddBodyFlag)
			case "multi-update":
				endpoint = c.MultiUpdate(storageMultiUpdateEncoderFn)
				data, err = storagec.BuildMultiUpdatePayload(*storageMultiUpdateBodyFlag, *storageMultiUpdateIdsFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// sommelierUsage displays the usage of the sommelier command and its
// subcommands.
func sommelierUsage() {
	fmt.Fprintf(os.Stderr, `The sommelier service retrieves bottles given a set of criteria.
Usage:
    %[1]s [globalflags] sommelier COMMAND [flags]

COMMAND:
    pick: Pick implements pick.

Additional help:
    %[1]s sommelier COMMAND --help
`, os.Args[0])
}
func sommelierPickUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] sommelier pick -body JSON

Pick implements pick.
    -body JSON: 

Example:
    %[1]s sommelier pick --body '{
      "name": "Blue\'s Cuvee",
      "varietal": [
         "pinot noir",
         "merlot",
         "cabernet franc"
      ],
      "winery": "longoria"
   }'
`, os.Args[0])
}

// storageUsage displays the usage of the storage command and its subcommands.
func storageUsage() {
	fmt.Fprintf(os.Stderr, `The storage service makes it possible to view, add or remove wine bottles.
Usage:
    %[1]s [globalflags] storage COMMAND [flags]

COMMAND:
    list: List all stored bottles
    show: Show bottle by ID
    add: Add new bottle and return its ID.
    remove: Remove bottle from storage
    rate: Rate bottles by IDs
    multi-add: Add n number of bottles and return their IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be added.
    multi-update: Update bottles with the given IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be updated. The IDs in the query parameter is mapped to each part in the request.

Additional help:
    %[1]s storage COMMAND --help
`, os.Args[0])
}
func storageListUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage list

List all stored bottles

Example:
    %[1]s storage list
`, os.Args[0])
}

func storageShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage show -id STRING -view STRING

Show bottle by ID
    -id STRING: ID of bottle to show
    -view STRING: 

Example:
    %[1]s storage show --id "Sapiente et." --view "tiny"
`, os.Args[0])
}

func storageAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage add -body JSON

Add new bottle and return its ID.
    -body JSON: 

Example:
    %[1]s storage add --body '{
      "composition": [
         {
            "percentage": 92,
            "varietal": "Syrah"
         },
         {
            "percentage": 92,
            "varietal": "Syrah"
         },
         {
            "percentage": 92,
            "varietal": "Syrah"
         },
         {
            "percentage": 92,
            "varietal": "Syrah"
         }
      ],
      "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
      "name": "Blue\'s Cuvee",
      "rating": 1,
      "vintage": 1941,
      "winery": {
         "country": "USA",
         "name": "Longoria",
         "region": "Central Coast, California",
         "url": "http://www.longoriawine.com/"
      }
   }'
`, os.Args[0])
}

func storageRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage remove -id STRING

Remove bottle from storage
    -id STRING: ID of bottle to remove

Example:
    %[1]s storage remove --id "Corporis rem."
`, os.Args[0])
}

func storageRateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage rate -p JSON

Rate bottles by IDs
    -p JSON: map[uint32][]string is the payload type of the storage service rate method.

Example:
    %[1]s storage rate --p '{
      "\u003cuint32 Value\u003e": [
         "Laboriosam consequatur delectus doloribus.",
         "Est mollitia.",
         "Voluptas ex enim.",
         "Est explicabo eveniet dolore."
      ]
   }'
`, os.Args[0])
}

func storageMultiAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage multi-add -body JSON

Add n number of bottles and return their IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be added.
    -body JSON: 

Example:
    %[1]s storage multi-add --body '[
      {
         "composition": [
            {
               "percentage": 92,
               "varietal": "Syrah"
            },
            {
               "percentage": 92,
               "varietal": "Syrah"
            }
         ],
         "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
         "name": "Blue\'s Cuvee",
         "rating": 1,
         "vintage": 1984,
         "winery": {
            "country": "USA",
            "name": "Longoria",
            "region": "Central Coast, California",
            "url": "http://www.longoriawine.com/"
         }
      },
      {
         "composition": [
            {
               "percentage": 92,
               "varietal": "Syrah"
            },
            {
               "percentage": 92,
               "varietal": "Syrah"
            }
         ],
         "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
         "name": "Blue\'s Cuvee",
         "rating": 1,
         "vintage": 1984,
         "winery": {
            "country": "USA",
            "name": "Longoria",
            "region": "Central Coast, California",
            "url": "http://www.longoriawine.com/"
         }
      }
   ]'
`, os.Args[0])
}

func storageMultiUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] storage multi-update -body JSON -ids JSON

Update bottles with the given IDs. This is a multipart request and each part has field name 'bottle' and contains the encoded bottle info to be updated. The IDs in the query parameter is mapped to each part in the request.
    -body JSON: 
    -ids JSON: 

Example:
    %[1]s storage multi-update --body '{
      "bottles": [
         {
            "composition": [
               {
                  "percentage": 92,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 92,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 1,
            "vintage": 1984,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 92,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 92,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 1,
            "vintage": 1984,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 92,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 92,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 1,
            "vintage": 1984,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         },
         {
            "composition": [
               {
                  "percentage": 92,
                  "varietal": "Syrah"
               },
               {
                  "percentage": 92,
                  "varietal": "Syrah"
               }
            ],
            "description": "Red wine blend with an emphasis on the Cabernet Franc grape and including other Bordeaux grape varietals and some Syrah",
            "name": "Blue\'s Cuvee",
            "rating": 1,
            "vintage": 1984,
            "winery": {
               "country": "USA",
               "name": "Longoria",
               "region": "Central Coast, California",
               "url": "http://www.longoriawine.com/"
            }
         }
      ]
   }' --ids '[
      "Aut rem vel veritatis.",
      "Animi nulla aut aut."
   ]'
`, os.Args[0])
}
