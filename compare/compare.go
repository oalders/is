// package compare compares versions
package compare

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/go-version"
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

func Strings(operator, got, want string) (bool, error) {
	var err error
	var success bool

	switch operator {
	case "like":
		success, err = regexp.MatchString(want, got)
	case "unlike":
		success, err = regexp.MatchString(want, got)
		if err == nil {
			success = !success
		} else {
			success = false
		}
	default:
		err = fmt.Errorf(
			"%s is not a string comparison operator",
			operator,
		)
	}
	return success, err
}
