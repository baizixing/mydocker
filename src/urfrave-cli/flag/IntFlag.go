package flag

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func IntFlag() {
	var number int
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "number",
				Usage:       "count the people numeber",
				Value:       0,
				Destination: &number,
				Action: func(ctx *cli.Context, v int) error {
					// if arg, _ := strconv.Atoi(ctx.Args().Get(0)); arg > 10 {
					// 	fmt.Println("input must be in range [0-10]!")
					// } else {
					// 	fmt.Println(ctx.Args().Get(0))
					// }
					// return nil

					if v > 10 {
						return fmt.Errorf("input must be in range [0-10]!")
					}
					return nil

				},
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Println(number)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
