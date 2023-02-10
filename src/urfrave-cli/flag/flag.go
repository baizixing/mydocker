package flag

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func flag() {
	var language string
	// var count bool
	app := &cli.App{
		Flags: []cli.Flag{

			&cli.StringFlag{
				Name:        "lang",
				Aliases:     []string{"l"},
				Value:       "english",
				Usage:       "language for the greeting",
				Destination: &language,
			},

			// &cli.BoolFlag{
			// 	Name:        "foo",
			// 	Usage:       "foo greeting",
			// 	// Destination: &count,
			// },
		},
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() <= 0 {
				fmt.Println("there must be something wrong with your input")
				return nil
			}

			for index := 0; index < ctx.NArg(); index++ {

				// if ctx.String("lang") == "spanish" {
				// 	fmt.Printf("hola %s \n", ctx.Args().Get(index))
				// } else {
				// 	fmt.Printf("hello %s \n", ctx.Args().Get(index))
				// }

				//使用destination来读取lang值的方案
				if language == "spanish" {
					fmt.Printf("hola %s \n", ctx.Args().Get(index))
				} else {
					fmt.Printf("hello %s \n", ctx.Args().Get(index))
				}

				// fmt.Println(count)

			}
			return nil
		},
	}

	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
