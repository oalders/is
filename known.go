// package main contains the logic for the "known" command
package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"runtime"
	"strings"

	"github.com/oalders/is/attr"
	"github.com/oalders/is/os"
	"github.com/oalders/is/parser"
	"github.com/oalders/is/types"
	"github.com/oalders/is/version"
)

// Run "is known ...".
//
//nolint:gocritic
func (r *KnownCmd) Run(ctx *types.Context) error {
	result := ""
	var err error

	isVersion, segment, err := isVersion(r)
	if err != nil {
		return err
	}

	if r.OS.Attr != "" {
		result, err = os.Info(ctx, r.OS.Attr)
	} else if r.CLI.Attr != "" {
		result, err = runCLI(ctx, r.CLI.Name)
	} else if r.Arch.Attr != "" {
		result = runtime.GOARCH
	}
	if err != nil {
		return err
	}

	if len(result) > 0 && isVersion {
		got, err := version.NewVersion(result)
		if err != nil {
			return fmt.Errorf("parse version from output: %w", err)
		}
		segments := got.Segments()
		result = fmt.Sprintf("%d", segments[segment])
	}

	if len(result) > 0 {
		ctx.Success = true
	}

	//nolint:forbidigo
	fmt.Println(result)

	return err
}

func isVersion(r *KnownCmd) (bool, uint, error) { //nolint:varnamelen
	if r.OS.Attr == attr.Version || r.CLI.Attr == attr.Version {
		switch {
		case r.Major:
			return true, 0, nil
		case r.Minor:
			return true, 1, nil
		case r.Patch:
			return true, 2, nil
		}
	}
	if r.Major || r.Minor || r.Patch {
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
		if err != nil {
			result = strings.TrimRight(result, "\n")
		}
	}
	return result, err
}
