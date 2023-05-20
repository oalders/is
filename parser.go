// This file contains output parsers
package main

import (
	"os/exec"
	"regexp"
)

func cliOutput(cliName string) (string, error) {
	versionArg := map[string]string{
		"go":   "version",
		"tmux": "-V",
	}
	arg := "--version"
	if v, exists := versionArg[cliName]; exists {
		arg = v
	}

	o, err := exec.Command("command", cliName, arg).Output()
	if err != nil {
		o, err = exec.Command(cliName, arg).Output()
		if err != nil {
			return "", err
		}
	}

	return cliVersion(cliName, string(o)), nil
}

func cliVersion(cliName, output string) string {
	regexen := map[string]string{
		"go":   `go version go(\d+\.\d+\.\d+)\s`,
		"perl": `This is perl .* \((v\d+\.\d+\.\d+)\)`,
		"tmux": `tmux (.*)\b`,
	}
	re := regexp.MustCompile(cliName + `\s+(.*)\b`)
	if v, exists := regexen[cliName]; exists {
		re = regexp.MustCompile(v)
	}
	matches := re.FindAllStringSubmatch(output, -1)
	if len(matches) > 0 {
		output = matches[0][1]
	}
	return output
}
