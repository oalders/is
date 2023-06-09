// The main package is the command line runner
package main

import (
	"os"

	"github.com/alecthomas/kong"
)

func main() {
	ctx := kong.Parse(&api,
		kong.Vars{
			"version": "0.1.0",
		})
	runContext := Context{Debug: api.Debug}
	err := ctx.Run(&runContext)
	ctx.FatalIfErrorf(err)

	if runContext.Success {
		os.Exit(0)
	}
	os.Exit(1)
}
