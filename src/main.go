package main

import (
	"log"
	"os"

	"component"

	"github.com/urfave/cli/v2"
)

func main() {
	//创建cli app，本次总共两个command，run和init
	app := &cli.App{
		Name:  "mydocker",
		Usage: "for docker learning",
		Commands: []*cli.Command{
			&component.MydockerRunCommand,
			&component.MydockerInitCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
