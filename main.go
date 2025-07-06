package main

import (
	"fmt"
	"os"

	"github.com/hadian90/hyprwwhide/cmd"
	"github.com/hadian90/hyprwwhide/utils"

	"github.com/urfave/cli/v2"
)

func main() {
	utils.CheckIfMainFolderExist()
	app := &cli.App{
		Name:  "hyprwwhide",
		Usage: "A command-line tool to manage windows.",
		Commands: []*cli.Command{
			cmd.HideActiveCmd,
			cmd.HideAllCmd,
			cmd.RevealCmd,
			cmd.RevealAllCmd,
			cmd.ShowCmd,
			cmd.ShowAllCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
