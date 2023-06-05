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
		"ssh":  "-V",
		"tmux": "-V",
	}
	arg := "--version"
	if v, exists := versionArg[cliName]; exists {
		arg = v
	}

	cmd := exec.Command("command", cliName, arg)
	o, err := cmd.Output()

	// ssh -V doesn't print to STDOUT?
	if len(o) == 0 && err == nil {
		cmd = exec.Command("command", cliName, arg)
		o, err = cmd.CombinedOutput()
	}

	if err != nil {
		o, err = exec.Command(cliName, arg).Output()
		if err != nil {
			o, err = exec.Command("/bin/bash", "-c", cliName, arg).Output()
			if err != nil {
				fmt.Printf("err %+v\n", err)
				return "", err
			}
		}
	}

	return cliVersion(ctx, cliName, string(o)), nil
}

func cliVersion(ctx *Context, cliName, output string) string {
	regexen := map[string]string{
		"ansible": `ansible \[core (\d+\.\d+\.\d+)\]`,
		"git":     `git version (\d+\.\d+\.\d+)\s`,
		"go":      `go version go(\d+\.\d+\.\d+)\s`,
		"perl":    `This is perl .* \((v\d+\.\d+\.\d+)\)`,
		"plenv":   `plenv ([\d\w\-\.]*)\b`,
		"python":  `Python ([0-9.]*)\b`,
		"python3": `Python ([0-9.]*)\b`,
		"rg":      `ripgrep ([0-9.]*)\b`,
		"ruby":    `(\d+\.\d+\.[\d\w]+)\b`,
		"ssh":     `OpenSSH_([0-9a-z.]*)\b`,
		"tmux":    `tmux (.*)\b`,
		"tree":    `(v\d+\.\d+\.\d+)\b`,
		"vim":     `VIM - Vi IMproved (\d+\.\d+)\s`,
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
