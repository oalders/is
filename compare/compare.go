// package compare compares versions
package compare

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/oalders/is/types"
	"github.com/oalders/is/version"
)

func CLIVersions(ctx *types.Context, op, g, w string) error {
	var success bool
	switch op {
	case "like", "unlike":
		return Strings(ctx, op, g, w)
	}
	got, err := version.NewVersion(g)
	if err != nil {
		return err
	}
	want, err := version.NewVersion(w)
	if err != nil {
		return err
	}

	switch op {
	case "eq":
		success = got.Equal(want)
	case "ne":
		success = got.Compare(want) != 0
	case "lt":
		success = got.LessThan(want)
	case "lte":
		success = got.Compare(want) <= 0
	case "gt":
		success = got.GreaterThan(want)
	case "gte":
		success = got.Compare(want) >= 0
	}

	ctx.Success = success
	return nil
}

func Strings(ctx *types.Context, operator, got, want string) error {
	var err error
	var success bool

	comparison := fmt.Sprintf(`comparison "%s" %s "%s"`, want, operator, got)
	if ctx.Debug {
		log.Print(comparison)
	}
	switch operator {
	case "eq":
		success = got == want
	case "ne":
		success = got != want
	case "like", "unlike":
		success, err = regexp.MatchString(want, got)
	}

	if err != nil {
		err = errors.Join(fmt.Errorf("error in comparison: %s", comparison), err)
	} else if operator == "unlike" {
		success = !success
	}
	ctx.Success = success
	return err
}
