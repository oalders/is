// Package parser contains output parsers
package parser

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/oalders/is/types"
)

func CLIOutput(ctx *types.Context, cliName string) (string, error) {
	versionArg := map[string]string{
		"go":      "version",
		"lua":     "-v",
		"openssl": "version",
		"pihole":  "-v",
		"ssh":     "-V",
		"tmux":    "-V",
	}
	arg := "--version"
	if v, exists := versionArg[cliName]; exists {
		arg = v
	}

	args := []string{cliName, arg}
	if ctx.Debug {
		log.Printf("Running: %s %s\n", args[0], args[1])
	}
	cmd := exec.Command(cliName, arg)
	output, err := cmd.Output()
	if err != nil {
		return "", errors.Join(errors.New("could not get command output"), err)
	}

	// ssh -V doesn't print to STDOUT?
	if len(output) == 0 {
		if ctx.Debug {
			log.Printf("Running: %s %s and checking STDERR\n", args[0], args[1])
		}
		cmd = exec.Command(cliName, arg)
		output, err = cmd.CombinedOutput()
		if err != nil {
			return "", errors.Join(errors.New("could not find command output in STDIN or STDERR"), err)
		}
	}

	return CLIVersion(ctx, cliName, string(output)), nil
}

func CLIVersion(ctx *types.Context, cliName, output string) string {
	floatRegex := `[\d.]*`
	floatWithTrailingLetterRegex := `[\d.]*\w`
	intRegex := `\d*`
	vStringRegex := `v[\d.]*`
	vStringWithTrailingLetterRegex := `v[\d.]*\w`
	regexen := map[string]string{
		"ansible": fmt.Sprintf(`ansible \[core (%s)\b`, floatRegex),
		"bash":    fmt.Sprintf(`version (%s)\b`, floatRegex),
		"bat":     fmt.Sprintf(`bat (%s)\b`, floatRegex),
		"curl":    fmt.Sprintf(`curl (%s)\b`, floatRegex),
		"docker":  fmt.Sprintf(`version (%s),`, floatRegex),
		"gcc":     fmt.Sprintf(`clang version (%s)\b`, floatRegex),
		"git":     fmt.Sprintf(`git version (%s)\s`, floatRegex),
		"gh":      fmt.Sprintf(`gh version (%s)\b`, floatRegex),
		"go":      fmt.Sprintf(`go version go(%s)\s`, floatRegex),
		"jq":      fmt.Sprintf(`jq-(%s)\b`, floatRegex),
		"less":    fmt.Sprintf(`less (%s)\b`, intRegex),
		"lua":     fmt.Sprintf(`Lua (%s)\b`, floatRegex),
		"md5sum":  fmt.Sprintf(`md5sum \(GNU coreutils\) (%s)\b`, floatRegex),
		"perl":    fmt.Sprintf(`This is perl .* \((%s)\)\s`, vStringRegex),
		"openssl": fmt.Sprintf(`SSL (%s)\b`, floatWithTrailingLetterRegex),
		"pihole":  fmt.Sprintf(`Pi-hole version is (%s)`, vStringRegex),
		"plenv":   `plenv ([\d\w\-\.]*)\b`,
		"python":  fmt.Sprintf(`Python (%s)\b`, floatRegex),
		"python3": fmt.Sprintf(`Python (%s)\b`, floatRegex),
		"rg":      fmt.Sprintf(`ripgrep (%s)\b`, floatRegex),
		"ruby":    `ruby (\d+\.\d+\.[\d\w]+)\b`,
		"ssh":     `OpenSSH_([0-9a-z.]*)\b`,
		"tar":     fmt.Sprintf(`bsdtar (%s)\b`, floatRegex),
		"tmux":    fmt.Sprintf(`tmux (%s)\b`, floatWithTrailingLetterRegex),
		"tree":    fmt.Sprintf(`tree (%s)\b`, vStringWithTrailingLetterRegex),
		"trurl":   fmt.Sprintf(`trurl version (%s)\b`, floatRegex),
		"unzip":   fmt.Sprintf(`UnZip (%s)\b`, floatRegex),
		"vim":     fmt.Sprintf(`VIM - Vi IMproved (%s)\b`, floatRegex),
		"zsh":     fmt.Sprintf(`zsh (%s)\b`, floatRegex),
	}
	var versionRegex *regexp.Regexp
	hasNewLines := regexp.MustCompile("\n")
	if v, exists := regexen[cliName]; exists {
		versionRegex = regexp.MustCompile(v)
	} else if found := len(hasNewLines.FindAllStringIndex(output, -1)); found > 1 {
		// If --version returns more than one line, the actual version will
		// generally be the last thing on the first line
		versionRegex = regexp.MustCompile(fmt.Sprintf(`(?:\s)(%s|%s|%s|%s)\s*\n`, vStringWithTrailingLetterRegex, floatWithTrailingLetterRegex, vStringRegex, floatRegex))
	} else {
		versionRegex = regexp.MustCompile(`(?i)` + cliName + `\s+(.*)\b`)
	}

	matches := versionRegex.FindAllStringSubmatch(output, -1)
	if len(matches) > 0 {
		output = matches[0][1]
	} else if ctx.Debug {
		log.Printf("output \"%s\" does not match regex \"%s\"\n", output, versionRegex)
	}
	output = strings.TrimRight(output, "\n")

	return output
}
