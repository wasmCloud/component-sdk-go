package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"

	"go.bytecodealliance.org/cmd/wit-bindgen-go/cmd/generate"
)

func main() {
	cmd := generate.Command
	cmd.Name = "component-sdk-go-codegen"
	cmd.Flags = append(cmd.Flags,
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "print verbose logging messages",
		},
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"vv"},
			Usage:   "print debug logging messages",
		},
	)

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
