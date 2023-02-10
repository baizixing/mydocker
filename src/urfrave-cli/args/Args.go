package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			// for index := 0; index < cCtx.NArg(); index++ {
			// 	fmt.Printf("hello %s \n", cCtx.Args().Get(index))
			// }

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
