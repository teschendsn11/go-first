package cli

import (
	"fmt"
	"os"

	"teschendsn11/go-first/h"

	"github.com/urfave/cli"
)

func StartCLI() *cli.App {
	h.Log.Debug("cli Start")

	cliapp := cli.App{
		Name:  "dnd cli",
		Usage: "creates dnd stuff",
		Commands: []cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					fmt.Println(os.Getenv("DND_DB"))
					return nil
				},
			},
			{
				Name:    "levelup",
				Aliases: []string{"lu"},
				Usage:   "complete a task on the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("completed task: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}

	if err := cliapp.Run(os.Args); err != nil {
		h.Log.Error(err.Error())
	}

	return &cliapp
}
