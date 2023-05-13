package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Debug bool `help: turn on debugging statements`
	OS    struct {
		Name struct {
			Eq  string `arg required:"" enum:"eq,ne"`
			Val string `arg:"" required:""`
		} `cmd`
	} `cmd:"" help:"Match on OS."`
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "rm <path>":
	case "os name <eq> <val>":
		// fmt.Printf("%+v", ctx.Args)
		expected := ctx.Args[len(ctx.Args)-1]
		if ctx.Args[2] == "eq" {
			if runtime.GOOS == expected {
				os.Exit(0)
			}
			fmt.Printf("Comparison failed: %s eq %s\n", runtime.GOOS, expected)
			os.Exit(1)
		}
		if ctx.Args[2] == "ne" {
			if ctx.Args[3] != runtime.GOOS {
				os.Exit(0)
			}
			fmt.Printf("Comparison failed: %s ne %s\n", runtime.GOOS, expected)
			os.Exit(1)
		}
		os.Exit(1)
	case "ls":
	default:
		panic(ctx.Command())
	}
}
