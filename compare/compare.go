// package compare compares versions
package compare

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/oalders/is/ops"
	"github.com/oalders/is/os"
	"github.com/oalders/is/types"
	"github.com/oalders/is/version"
)

func CLIVersions(ctx *types.Context, operator, gotString, wantString string) error {
	var success bool
	switch operator {
	case ops.Like, ops.Unlike:
		return Strings(ctx, operator, gotString, wantString)
	}
	got, err := version.NewVersion(gotString)
	if err != nil {
		return err
	}
	want, err := version.NewVersion(wantString)
	if err != nil {
		return err
	}

	switch operator {
	case ops.Eq:
		success = got.Equal(want)
	case ops.Ne:
		success = got.Compare(want) != 0
	case ops.Lt:
		success = got.LessThan(want)
	case ops.Lte:
		success = got.Compare(want) <= 0
	case ops.Gt:
		success = got.GreaterThan(want)
	case ops.Gte:
		success = got.Compare(want) >= 0
	}

	ctx.Success = success
	return maybeDebug(ctx, operator, gotString, wantString)
}

func Strings(ctx *types.Context, operator, got, want string) error {
	var err error
	var success bool

	comparison := fmt.Sprintf(`comparison "%s" %s "%s"`, want, operator, got)
	if ctx.Debug {
		log.Print(comparison)
	}
	switch operator {
	case ops.Eq:
		success = got == want
	case ops.Ne:
		success = got != want
	case ops.Like, ops.Unlike:
		success, err = regexp.MatchString(want, got)
	}

	if err != nil {
		ctx.Success = false
		return errors.Join(fmt.Errorf("error in comparison: %s", comparison), err)
	}
	if operator == ops.Unlike {
		success = !success
	}
	ctx.Success = success
	return maybeDebug(ctx, operator, got, want)
}

func maybeDebug(ctx *types.Context, operator, got, want string) error {
	if !ctx.Debug {
		return nil
	}

	if !ctx.Success {
		log.Printf("Comparison failed: %s %s %s\n", got, operator, want)
	}
	os, err := os.Aggregated(ctx)
	if err != nil {
		return err
	}
	log.Printf("%s\n", os)
	return nil
}
