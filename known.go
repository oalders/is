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

	if result != "" {
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
		result, err := is_os.Aggregated(ctx)
		if err != nil {
			return err
		}
		success(ctx, result)
		return nil
	}
	if attr == "battery" {
		batt, err := battery.Get(ctx, nth)
		if err != nil {
			return err
		}
		if asJSON {
			data, err := json.MarshalIndent(batt, "", "    ")
			if err != nil {
				return errors.Join(
					fmt.Errorf("could not marshal indented JSON (%+v)", batt),
					err,
				)
			}
			success(ctx, string(data))
			return nil
		}

		headers := []string{
			"Attribute",
			"Value",
			"Units",
		}

		rows := [][]string{
			{"battery-number", fmt.Sprintf("%d", batt.BatteryNumber)},
			{"charge-rate", fmt.Sprintf("%v", batt.ChargeRate), "mW"},
			{"count", fmt.Sprintf("%d", batt.Count)},
			{"current-capacity", fmt.Sprintf("%v", batt.CurrentCapacity), "mWh"},
			{"current-charge", fmt.Sprintf("%v", batt.CurrentCharge), "%"},
			{"design-capacity", fmt.Sprintf("%v", batt.DesignCapacity), "mWh"},
			{"design-voltage", fmt.Sprintf("%v", batt.DesignVoltage), "mWh"},
			{"last-full-capacity", fmt.Sprintf("%v", batt.LastFullCapacity), "mWh"},
			{"state", fmt.Sprintf("%v", batt.State), ""},
			{"voltage", fmt.Sprintf("%v", batt.Voltage), "V"},
		}
		success(ctx, tabular(headers, rows))
		return nil
	}
	return fmt.Errorf("unknown attribute: %s", attr)
}
