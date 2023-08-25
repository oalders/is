// package main contains macOS logic
package mac

import (
	"os/exec"
	"strings"

	"github.com/hashicorp/go-version"
)

func CodeName(osVersion string) string {
	got, err := version.NewVersion(osVersion)
	if err != nil {
		return ""
	}
	// https://en.wikipedia.org/wiki/List_of_Apple_codenames
	segments := got.Segments()
	name := ""
	switch segments[0] {
	case 13:
		name = "ventura"
	case 12:
		name = "monterey"
	case 11:
		name = "big sur"
	case 10:
		switch segments[1] {
		case 15:
			name = "catalina"
		case 14:
			name = "mojave"
		case 13:
			name = "high sierra"
		case 12:
			name = "sierra"
		case 11:
			name = "el capitan"
		case 10:
			name = "yosemite"
		case 9:
			name = "mavericks"
		case 8:
			name = "mountain lion" // released 2012
		}
	}

	return name
}

func Version() (string, error) {
	o, err := exec.Command("sw_vers", "-productVersion").Output()
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(o), "\n"), nil
}
