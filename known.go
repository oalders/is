// package main contains the logic for the "known" command
package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/oalders/is/types"
)

// Run "is known ..."
func (r *KnownCmd) Run(ctx *types.Context) error {
	result := ""
	var err error

	if r.OS.Attr != "" {
		result, err = osInfo(ctx, r.OS.Attr)
		if ctx.Debug {
			os, err := aggregatedOS()
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", os)
		}
	} else if r.CLI.Attr != "" {
		result, err = cliOutput(ctx, r.CLI.Name)
		if err != nil {
			re := regexp.MustCompile(`executable file not found`)
			if re.MatchString(err.Error()) {
				if ctx.Debug {
					fmt.Printf("executable file \"%s\" not found", r.CLI.Name)
				}

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
	fmt.Println(result)
	ctx.Success = true

	return nil
}
