// package main contains the logic for the "known" command
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/oalders/is/attr"
	"github.com/oalders/is/battery"
	is_os "github.com/oalders/is/os"
	"github.com/oalders/is/parser"
	"github.com/oalders/is/types"
	"github.com/oalders/is/version"
)

const (
	manpath = "MANPATH"
	path    = "PATH"
)

// Run "is known ...".
//
//nolint:cyclop
func (r *KnownCmd) Run(ctx *types.Context) error {
	result := ""
	{
		var err error

		switch {
		case r.Summary.Attr != "":
			return summary(ctx, r.Summary.Attr, r.Summary.Nth, r.Summary.JSON)
		case r.OS.Attr != "":
			result, err = is_os.Info(ctx, r.OS.Attr)
		case r.CLI.Attr != "":
			result, err = runCLI(ctx, r.CLI.Name)
		case r.Var.Name != "":
			result, err = getEnv(ctx, r.Var.Name, r.Var.JSON)
		case r.Battery.Attr != "":
			result, err = battery.GetAttrAsString(
				ctx,
				r.Battery.Attr,
				r.Battery.Round,
				r.Battery.Nth,
			)
		case r.Arch.Attr != "":
			success(ctx, runtime.GOARCH)
			return nil
		}
		if err != nil {
			return err
		}
	}

	if result != "" {
		isVersion, segment, versionErr := isVersion(r)
		if versionErr != nil {
			return versionErr
		}

		if isVersion {
			got, versionErr := version.NewVersion(result)
			if versionErr != nil {
				return fmt.Errorf("parse version from output: %w", versionErr)
			}
			segments := got.Segments()
			result = fmt.Sprintf("%d", segments[segment])
		}
	}

	if r.Var.Name == "" && result != "" {
		ctx.Success = true
	}

	//nolint:forbidigo
	fmt.Println(result)
	return nil
}

//nolint:cyclop
func isVersion(r *KnownCmd) (bool, uint, error) { //nolint:varnamelen
	if r.OS.Attr == attr.Version || r.CLI.Attr == attr.Version {
		switch {
		case r.OS.Major || r.CLI.Major:
			return true, 0, nil
		case r.OS.Minor || r.CLI.Minor:
			return true, 1, nil
		case r.OS.Patch || r.CLI.Patch:
			return true, 2, nil
		}
	}
	if r.OS.Major || r.OS.Minor || r.OS.Patch || r.CLI.Major || r.CLI.Minor || r.CLI.Patch {
		return false, 0, errors.New("--major, --minor and --patch can only be used with version")
	}
	return false, 0, nil
}

func runCLI(ctx *types.Context, cliName string) (string, error) {
	result, err := parser.CLIOutput(ctx, cliName)
	if err != nil {
		re := regexp.MustCompile(`executable file not found`)
		if re.MatchString(err.Error()) {
			if ctx.Debug {
				log.Printf("executable file \"%s\" not found", cliName)
			}

			ctx.Success = false
			return "", nil
		}

		return "", err
	}
	if len(result) > 0 {
		result = strings.TrimRight(result, "\n")
	}
	return result, err
}

func tabular(headers []string, rows [][]string) string {
	renderer := lipgloss.NewRenderer(os.Stdout)

	return table.New().
		Headers(headers...).
		Rows(rows...).
		Border(lipgloss.ThickBorder()).
		BorderStyle(renderer.NewStyle().Foreground(lipgloss.Color("238"))).
		BorderRow(true).
		StyleFunc(func(_, _ int) lipgloss.Style {
			return renderer.NewStyle().Padding(0, 1)
		}).String()
}

func success(ctx *types.Context, msg string) {
	fmt.Println(msg) //nolint:forbidigo
	ctx.Success = true
}

func summary(ctx *types.Context, attr string, nth int, asJSON bool) error {
	if attr == "os" {
		return osSummary(ctx, asJSON)
	}
	if attr == "battery" {
		return batterySummary(ctx, nth, asJSON)
	}
	if attr == "var" {
		return envSummary(ctx, asJSON)
	}
	return fmt.Errorf("unknown attribute: %s", attr)
}

func toJSON(record any) (string, error) {
	data, err := json.MarshalIndent(record, "", "    ")
	if err != nil {
		return "", fmt.Errorf("could not marshal indented JSON (%+v): %w", record, err)
	}

	return string(data), nil
}

func osSummary(ctx *types.Context, asJSON bool) error {
	summary, err := is_os.ReleaseSummary(ctx)
	if err != nil {
		return err
	}
	if asJSON {
		result, err := toJSON(summary)
		if err != nil {
			return err
		}
		success(ctx, result)
		return nil
	}
	headers := []string{
		"Attribute",
		"Value",
	}

	rows := [][]string{
		{"name", summary.Name},
		{"version", summary.Version},
		{"version-codename", summary.VersionCodeName},
	}

	if summary.ID != "" {
		rows = append(rows, []string{"id", summary.ID})
	}
	if summary.IDLike != "" {
		rows = append(rows, []string{"id-like", summary.IDLike})
	}
	if summary.PrettyName != "" {
		rows = append(rows, []string{"pretty-name", summary.PrettyName})
	}
	success(ctx, tabular(headers, rows))
	return nil
}

func batterySummary(ctx *types.Context, nth int, asJSON bool) error {
	summary, err := battery.Get(ctx, nth)
	if err != nil {
		return err
	}
	if asJSON {
		if summary.Count == 0 {
			result, err := toJSON(map[string]int{"count": 0})
			if err != nil {
				return err
			}
			success(ctx, result)
			return nil
		}
		result, err := toJSON(summary)
		if err != nil {
			return err
		}
		success(ctx, result)
		return nil
	}

	headers := []string{
		"Attribute",
		"Value",
	}

	var rows [][]string

	if summary.Count > 0 {
		rows = [][]string{
			{"battery-number", fmt.Sprintf("%d", summary.BatteryNumber)},
			{"charge-rate", fmt.Sprintf("%v mW", summary.ChargeRate)},
			{"count", fmt.Sprintf("%d", summary.Count)},
			{"current-capacity", fmt.Sprintf("%v mWh", summary.CurrentCapacity)},
			{"current-charge", fmt.Sprintf("%v %%", summary.CurrentCharge)},
			{"design-capacity", fmt.Sprintf("%v mWh", summary.DesignCapacity)},
			{"design-voltage", fmt.Sprintf("%v mWh", summary.DesignVoltage)},
			{"last-full-capacity", fmt.Sprintf("%v mWh", summary.LastFullCapacity)},
			{"state", fmt.Sprintf("%v", summary.State)},
			{"voltage", fmt.Sprintf("%v V", summary.Voltage)},
		}
	} else {
		rows = append(rows, []string{"count", "0"})
	}
	success(ctx, tabular(headers, rows))
	return nil
}

func envSummary(ctx *types.Context, asJSON bool) error {
	envMap := make(map[string]any)
	rows := make([][]string, 0, len(os.Environ()))

	envVars := os.Environ()
	sort.Strings(envVars)

	for _, entry := range envVars {
		parts := strings.SplitN(entry, "=", 2)
		if len(parts) != 2 {
			continue
		}

		name, value := parts[0], parts[1]

		if name == path || name == manpath {
			pathParts := strings.Split(value, ":")
			if asJSON {
				envMap[name] = pathParts
			} else {
				rows = append(rows, []string{name, strings.Join(pathParts, "\n")})
			}
			continue
		}

		// Handle regular environment variables
		if asJSON {
			envMap[name] = value
		} else {
			rows = append(rows, []string{name, value})
		}
	}

	if asJSON {
		result, err := toJSON(envMap)
		if err != nil {
			return err
		}
		success(ctx, result)
		return nil
	}

	// Tabular output
	headers := []string{"Name", "Value"}
	success(ctx, tabular(headers, rows))
	return nil
}

func getEnv(ctx *types.Context, name string, asJSON bool) (string, error) {
	value, set := os.LookupEnv(name)
	ctx.Success = set

	if !asJSON {
		return value, nil
	}

	values := []string{}

	switch {
	case set && (name == path || name == manpath):
		values = strings.Split(value, ":")
	case set:
		values = append(values, value)
	default:
		values = nil
	}

	result, err := toJSON(values)
	if err != nil {
		return "", err
	}

	return result, nil
}
