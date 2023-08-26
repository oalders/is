// package compare compares versions
package compare

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/hashicorp/go-version"
	"github.com/oalders/is/types"
)

func CLIVersions(op string, got, want *version.Version) bool {
	var success bool
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

	return success
}

func Strings(ctx *types.Context, operator, got, want string) error {
	var err error
	var success bool

	if ctx.Debug {
		log.Printf(`comparing regex "%s" with %s`+"\n", want, got)
	}
	switch operator {
	case "like", "unlike":
		success, err = regexp.MatchString(want, got)
	default:
		err = fmt.Errorf(
			"%s is not a string comparison operator",
			operator,
		)
	}

	switch operator {
	case "like":
		if err != nil {
			err = errors.Join(fmt.Errorf("could not compare the version (%s) using (%s)", got, want), err)
		}
	case "unlike":
		if err != nil {
			err = errors.Join(fmt.Errorf("could not compare the version (%s) using (%s)", got, want), err)
			success = false
		} else {
			success = !success
		}
	}
	ctx.Success = success
	return err
}
