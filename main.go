// The main package is the command line runner for the is command.
package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/oalders/is/types"
	"github.com/posener/complete"
	"github.com/willabides/kongplete"
)

func main() {
	//nolint:lll,govet,nolintlint
	var API struct {
		Arch    ArchCmd          `cmd:"" help:"Check arch e.g. \"is arch like x64\""`
		Battery BatteryCmd       `cmd:"" help:"Check battery attributes. e.g. \"is battery state eq charging\""`
		CLI     CLICmd           `cmd:"" help:"Check cli version. e.g. \"is cli version tmux gte 3\""`
		Debug   bool             `help:"turn on debugging statements"`
		FSO     FSOCmd           `cmd:"" help:"Check fso (file system object). e.g. \"is fso age gte 3 days\""`
		Known   KnownCmd         `cmd:""`
		OS      OSCmd            `cmd:"" help:"Check OS attributes. e.g. \"is os name eq darwin\""`
		There   ThereCmd         `cmd:"" help:"Check if command exists. e.g. \"is there git\""`
		User    UserCmd          `cmd:"" help:"Info about current user. e.g. \"is user sudoer\""`
		Var     VarCmd           `cmd:"" help:"Check environment variables. e.g. \"is var EDITOR eq nvim\""`
		Version kong.VersionFlag `help:"Print version to screen"`

		InstallCompletions kongplete.InstallCompletions `cmd:"" help:"install shell completions. e.g. \"is install-completions\" and then run the command which is printed to your terminal to get completion in your current session. add the command to a .bashrc or similar to get completion across all sessions."` //nolint:lll
	}

	parser := kong.Must(&API,
		kong.Name("is"),
		kong.Description("an inspector for your environment"),
		kong.UsageOnError(),
		kong.Vars{"version": "0.8.2"},
	)

	// Run kongplete.Complete to handle completion requests
	kongplete.Complete(parser,
		kongplete.WithPredictor("file", complete.PredictFiles("*")),
	)

	ctx, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	runContext := types.Context{Debug: API.Debug}
	err = ctx.Run(&runContext)
	ctx.FatalIfErrorf(err)

	if runContext.Success {
		os.Exit(0)
	}
	os.Exit(1)
}
