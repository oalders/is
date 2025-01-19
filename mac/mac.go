// package main contains macOS logic
package mac

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/oalders/is/version"
)

// TODO: return error when no code name found.
//
//nolint:godox
func CodeName(osVersion string) string {
	got, err := version.NewVersion(osVersion)
	if err != nil {
		return ""
	}

	// https://en.wikipedia.org/wiki/List_of_Apple_codenames
	major := map[int]string{
		15: "sequoia",
		14: "sonoma",
		13: "ventura",
		12: "monterey",
		11: "big sur",
	}

	segments := got.Segments()

	if v, ok := major[segments[0]]; ok {
		return v
	}

	if segments[0] != 10 {
		return ""
	}

	minor := map[int]string{
		15: "catalina",
		14: "mojave",
		13: "high sierra",
		12: "sierra",
		11: "el capitan",
		10: "yosemite",
		9:  "mavericks",
		8:  "mountain lion", // released 2012
	}

	if v, ok := minor[segments[1]]; ok {
		return v
	}

	return ""
}

func Version() (string, error) {
	o, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		return "", errors.Join(errors.New("could not run sw_vers -productVersion"), err)
	}

	return strings.TrimRight(string(o), "\n"), nil
}
