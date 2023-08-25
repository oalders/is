// The main package is the command line runner
package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/oalders/is/types"
)

func main() {
	ctx := kong.Parse(&api,
		kong.Vars{
			"version": "0.1.1",
		})
	runContext := types.Context{Debug: api.Debug}
	err := ctx.Run(&runContext)
	ctx.FatalIfErrorf(err)

	if runContext.Success {
		os.Exit(0)
	}
	os.Exit(1)
}
