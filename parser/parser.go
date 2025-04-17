// Package parser contains output parsers
package parser

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/oalders/is/types"
)

func CLIOutput(ctx *types.Context, cliName string) (string, error) {
	versionArg := map[string]string{
		"dig":     "-v",
		"hugo":    "version",
		"go":      "version",
		"lua":     "-v",
		"openssl": "version",
		"perldoc": "-V",
		"pihole":  "-v",
		"ssh":     "-V",
		"tmux":    "-V",
	}
	arg := "--version"

	baseName := filepath.Base(cliName) // might be a path
	if v, exists := versionArg[baseName]; exists {
		arg = v
	}

	args := []string{cliName, arg}
	if ctx.Debug {
		log.Printf("Running: %s %s\n", args[0], args[1])
	}
	cmd := exec.Command(cliName, arg)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("command output: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("error output: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("starting command: %w", err)
	}

	output, _ := io.ReadAll(stdout)
	// ssh -V doesn't print to STDOUT?
	if len(output) == 0 {
		if ctx.Debug {
			log.Printf("Running: %s %s and checking STDERR\n", args[0], args[1])
		}

		output, _ = io.ReadAll(stderr)
	}

	return CLIVersion(ctx, baseName, string(output)), nil
}

//nolint:funlen
func CLIVersion(ctx *types.Context, cliName, output string) string {
	floatRegex := `\d+\.\d+`
	floatWithTrailingLetterRegex := `[\d.]*\w`
	intRegex := `\d*`
	optimisticRegex := `[\d.]*`
	semverRegex := `\d+\.\d+\.\d+`
	vStringRegex := `v[\d.]*`
	vStringWithTrailingLetterRegex := `v[\d.]*\w`
	vStringWithTrailingGreedyRegex := `v[\d.]*[\w+-]*\w`
	regexen := map[string]string{
		"ansible":       fmt.Sprintf(`ansible \[core (%s)\b`, semverRegex),
		"bash":          fmt.Sprintf(`version (%s)\b`, semverRegex),
		"bat":           fmt.Sprintf(`bat (%s)\b`, semverRegex),
		"csh":           fmt.Sprintf(`(%s)`, semverRegex),
		"curl":          fmt.Sprintf(`curl (%s)\b`, semverRegex),
		"docker":        fmt.Sprintf(`version (%s),`, semverRegex),
		"fpp":           fmt.Sprintf(`version (%s)\b`, semverRegex),
		"fzf":           fmt.Sprintf(`(%s)\b`, semverRegex),
		"gcc":           fmt.Sprintf(`clang version (%s)\b`, semverRegex),
		"git":           fmt.Sprintf(`git version (%s)\s`, semverRegex),
		"gh":            fmt.Sprintf(`gh version (%s)\b`, semverRegex),
		"go":            fmt.Sprintf(`go version go(%s)\s`, semverRegex),
		"golangci-lint": fmt.Sprintf(`golangci-lint has version (%s)\s`, semverRegex),
		"grep":          fmt.Sprintf(`(%s|%s)`, semverRegex, floatRegex),
		"hugo":          fmt.Sprintf(`hugo v(%s)\b`, semverRegex),
		"jq":            fmt.Sprintf(`jq-(%s)\b`, floatRegex),
		"less":          fmt.Sprintf(`less (%s)\b`, intRegex),
		"lua":           fmt.Sprintf(`Lua (%s)\b`, semverRegex),
		"md5sum":        fmt.Sprintf(`md5sum \(GNU coreutils\) (%s)\b`, floatRegex),
		"nvim":          fmt.Sprintf(`NVIM (%s)\b`, vStringWithTrailingGreedyRegex),
		"perl":          fmt.Sprintf(`This is perl .* \((%s)\)\s`, vStringRegex),
		"ocaml":         fmt.Sprintf(`The OCaml toplevel, version (%s)`, semverRegex),
		"opam":          fmt.Sprintf(`(%s)`, semverRegex),
		"openssl":       fmt.Sprintf(`SSL (%s)\b`, floatWithTrailingLetterRegex),
		"perldoc":       fmt.Sprintf(`(%s)\b`, vStringRegex),
		"pihole":        fmt.Sprintf(`Pi-hole version is (%s)`, vStringRegex),
		"plenv":         `plenv ([\d\w\-\.]*)\b`,
		"python":        fmt.Sprintf(`Python (%s)\b`, semverRegex),
		"python3":       fmt.Sprintf(`Python (%s)\b`, semverRegex),
		"rg":            fmt.Sprintf(`ripgrep (%s)\b`, semverRegex),
		"ruby":          `ruby (\d+\.\d+\.[\d\w]+)\b`,
		"tcsh":          fmt.Sprintf(`(%s)`, semverRegex),
		"rustc":         fmt.Sprintf(`rustc (%s)\b`, semverRegex),
		"screen":        fmt.Sprintf(`version (%s)\b`, semverRegex),
		"sh":            fmt.Sprintf(`version (%s)\b`, semverRegex),
		"sqlite3":       fmt.Sprintf(`(%s)\b`, semverRegex),
		"ssh":           `OpenSSH_([0-9a-z.]*)\b`,
		"tar":           fmt.Sprintf(`bsdtar (%s)\b`, semverRegex),
		"typos":         fmt.Sprintf(`typos-cli (%s)\b`, semverRegex),
		"tmux":          fmt.Sprintf(`tmux (%s)\b`, floatWithTrailingLetterRegex),
		"tree":          fmt.Sprintf(`tree (%s)\b`, vStringWithTrailingLetterRegex),
		"trurl":         fmt.Sprintf(`trurl version (%s)\b`, floatRegex),
		"unzip":         fmt.Sprintf(`UnZip (%s)\b`, floatRegex),
		"vim":           fmt.Sprintf(`VIM - Vi IMproved (%s)\b`, floatRegex),
		"zsh":           fmt.Sprintf(`zsh (%s)\b`, floatRegex),
	}
	var versionRegex *regexp.Regexp
	hasNewLines := regexp.MustCompile("\n")
	if v, exists := regexen[cliName]; exists {
		versionRegex = regexp.MustCompile(v)
	} else if found := len(hasNewLines.FindAllStringIndex(output, -1)); found > 1 {
		// If --version returns more than one line, the actual version will
		// generally be the last thing on the first line
		versionRegex = regexp.MustCompile(fmt.Sprintf(`(?:\s)(%s|%s|%s|%s|%s|%s)\s*\n`,
			semverRegex, vStringWithTrailingLetterRegex, floatWithTrailingLetterRegex,
			vStringRegex, optimisticRegex, floatRegex))
	} else {
		versionRegex = regexp.MustCompile(`(?i)` + cliName + `\s+(.*)\b`)
	}

	matches := versionRegex.FindAllStringSubmatch(output, -1)
	if ctx.Debug {
		log.Printf("matching output \"%s\" on regex \"%s\"\n", output, versionRegex)
	}
	if len(matches) > 0 {
		output = matches[0][1]
	}
	output = strings.TrimRight(output, "\n")

	return output
}
