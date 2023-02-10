package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	//创建cli app，本次总共两个command，run和init
	app := cli.App{
		Name:  "mydocker",
		Usage: "for docker learning",
		Commands: []*cli.Command{
			&mydockerRunCommand,
			&mydockerInitCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	fmt.Println("hello, world!")
}
