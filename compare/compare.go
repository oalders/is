// package compare compares versions
package compare

import "github.com/hashicorp/go-version"

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
