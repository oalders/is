// package main contains the logic for the "known" command
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"runtime"
	"strings"

	"github.com/oalders/is/attr"
	"github.com/oalders/is/battery"
	"github.com/oalders/is/os"
	"github.com/oalders/is/parser"
	"github.com/oalders/is/types"
	"github.com/oalders/is/version"
)

// Run "is known ...".
//
//nolint:cyclop
func (r *KnownCmd) Run(ctx *types.Context) error {
	result := ""
	{
		var err error

		switch {
		case r.OS.Attr != "":
			result, err = os.Info(ctx, r.OS.Attr)
		case r.CLI.Attr != "":
			result, err = runCLI(ctx, r.CLI.Name)
		case r.Battery.Attr != "":
			result, err = battery.GetAttrAsString(
				ctx,
				r.Battery.Attr,
				r.Battery.Round,
				r.Battery.Nth,
			)
		case r.Summary.Attr != "":
			if r.Summary.Attr == "os" {
				result, err = os.Aggregated(ctx)
				if err != nil {
					return err
				}
			} else if r.Summary.Attr == "battery" {
				batt, err := battery.Get(ctx, r.Summary.Nth)
				if err != nil {
					return err
				}
				data, err := json.MarshalIndent(batt, "", "    ")
				if err != nil {
					return errors.Join(
						fmt.Errorf("could not marshal indented JSON (%+v)", batt),
						err,
					)
				}
				result = string(data)
			}
		case r.Arch.Attr != "":
			result = runtime.GOARCH
		}
		if err != nil {
			return err
		}
	}

	if result != "" {
		isVersion, segment, versionErr := isVersion(r)
		if versionErr != nil {
			return fmt.Errorf("parse version from output: %w", versionErr)
		}
		segments := got.Segments()
		result = fmt.Sprintf("%d", segments[segment])
	}

	if result != "" {
		ctx.Success = true
	}

	//nolint:forbidigo
	fmt.Println(result)
	return nil
}

//nolint:cyclop
func isVersion(r *KnownCmd) (bool, uint, error) { //nolint:varnamelen
	if r.OS.Attr == attr.Version || r.CLI.Attr == attr.Version {
		switch {
		case r.OS.Major || r.CLI.Major:
			return true, 0, nil
		case r.OS.Minor || r.CLI.Minor:
			return true, 1, nil
		case r.OS.Patch || r.CLI.Patch:
			return true, 2, nil
		}
	}
	if r.OS.Major || r.OS.Minor || r.OS.Patch || r.CLI.Major || r.CLI.Minor || r.CLI.Patch {
		return false, 0, errors.New("--major, --minor and --patch can only be used with version")
	}
	return false, 0, nil
}

func runCLI(ctx *types.Context, cliName string) (string, error) {
	result, err := parser.CLIOutput(ctx, cliName)
	if err != nil {
		re := regexp.MustCompile(`executable file not found`)
		if re.MatchString(err.Error()) {
			if ctx.Debug {
				log.Printf("executable file \"%s\" not found", cliName)
			}

			ctx.Success = false
			return "", nil
		}

		return "", err
	}
	if len(result) > 0 {
		result = strings.TrimRight(result, "\n")
	}
	return result, err
}
