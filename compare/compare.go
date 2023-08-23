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

func Strings(op, got, want string) (bool, error) {
	var success bool

	switch op {
	case "like":
		return regexp.MatchString(want, got)
	case "unlike":
		match, err := regexp.MatchString(want, got)
		return !match, err
	}

	return success, fmt.Errorf(
		"%s is not a string comparison operator",
		op,
	)
}
