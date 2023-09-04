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
	} else if r.CLI.Attr != "" {
		result, err = runCLI(ctx, r.CLI.Name)
	}

	if err != nil {
		return err
	}
	if len(result) > 0 {
		ctx.Success = true
	}

	return err
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
		if err != nil {
			result = strings.TrimRight(result, "\n")
		}
	}
	return result, err
}
