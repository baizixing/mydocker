package flag

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func FilePaths() {
	var password string
	app := &cli.App{

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "database",
				Usage:       "use for database password",
				Aliases:     []string{"p"},
				Destination: &password,
				FilePath:    "password.txt",
				Required:    true,
			},
		},
		Action: func(ctx *cli.Context) error {
			if password == "111" {
				fmt.Println("pass!!!")
			} else {
				fmt.Println("password wrong!!!")
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
