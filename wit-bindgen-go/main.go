package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"

	"go.bytecodealliance.org/cmd/wit-bindgen-go/cmd/generate"
)

func main() {
	err := Command.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

var Command = &cli.Command{
	Name:  "wit-bindgen-go",
	Usage: "codegen helper",
	Commands: []*cli.Command{
		generate.Command,
	},
	Flags: []cli.Flag{
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
	},
	Action: func(ctx context.Context, cmd *cli.Command) error {
		return cli.ShowAppHelp(cmd)
	},
}
