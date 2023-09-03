// package main contains the logic for the "known" command
package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/oalders/is/os"
	"github.com/oalders/is/parser"
	"github.com/oalders/is/types"
)

// Run "is known ...".
func (r *KnownCmd) Run(ctx *types.Context) error {
	result := ""
	var err error

	if r.OS.Attr != "" {
		result, err = os.Info(ctx, r.OS.Attr)
		if ctx.Debug {
			os, err := os.Aggregated(ctx)
			if err != nil {
				return err
			}
			log.Printf("%s\n", os)
		}
	} else if r.CLI.Attr != "" {
		result, err = parser.CLIOutput(ctx, r.CLI.Name)
		if err != nil {
			re := regexp.MustCompile(`executable file not found`)
			if re.MatchString(err.Error()) {
				if ctx.Debug {
					log.Printf("executable file \"%s\" not found", r.CLI.Name)
				}

				ctx.Success = false
				return nil
			}

			return err
		}
		if len(result) > 0 {
			if err != nil {
				result = strings.TrimRight(result, "\n")
			}
		}
	}

	if err != nil {
		return err
	}
	if len(result) > 0 {
		ctx.Success = true
	}

	return err
}
