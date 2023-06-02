// This file contains output parsers
package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func cliOutput(ctx *Context, cliName string) (string, error) {
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

	return cliVersion(ctx, cliName, string(o)), nil
}

func cliVersion(ctx *Context, cliName, output string) string {
	regexen := map[string]string{
		"git":  `git version (\d+\.\d+\.\d+)\s`,
		"go":   `go version go(\d+\.\d+\.\d+)\s`,
		"perl": `This is perl .* \((v\d+\.\d+\.\d+)\)`,
		"tmux": `tmux (.*)\b`,
		"vim":  `VIM - Vi IMproved (\d+\.\d+)\s`,
	}
	var re *regexp.Regexp
	if v, exists := regexen[cliName]; exists {
		re = regexp.MustCompile(v)
	} else {
		re = regexp.MustCompile(`(?i)` + cliName + `\s+(.*)\b`)
	}

	matches := re.FindAllStringSubmatch(output, -1)
	if len(matches) > 0 {
		output = matches[0][1]
	} else if ctx.Debug {
		fmt.Printf("output \"%s\" does not match regex \"%s\"\n", output, re)
	}
	return output
}
