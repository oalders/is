// package compare compares versions
package compare

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/oalders/is/ops"
	"github.com/oalders/is/types"
	"github.com/oalders/is/version"
)

func CLIVersions(ctx *types.Context, operator, g, w string) error {
	var success bool
	switch operator {
	case ops.Like, ops.Unlike:
		return Strings(ctx, operator, g, w)
	}
	got, err := version.NewVersion(g)
	if err != nil {
		return err
	}
	want, err := version.NewVersion(w)
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
	case ops.Eq:
		success = got == want
	case ops.Ne:
		success = got != want
	case ops.Like, ops.Unlike:
		success, err = regexp.MatchString(want, got)
	}

	if err != nil {
		err = errors.Join(fmt.Errorf("error in comparison: %s", comparison), err)
	} else if operator == ops.Unlike {
		success = !success
	}
	ctx.Success = success
	return err
}
