package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ritarock/findinstance/lib/action"

	"github.com/urfave/cli"
)

func main() {
	var profile string

	app := cli.App{
		Name:  "findinstance",
		Usage: "Find ec2 instance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "profile",
				Usage:       "set aws profile",
				Value:       "default",
				Destination: &profile,
			},
		},
		Action: func(c *cli.Context) error {
			var instanceName string
			if c.NArg() > 0 {
				instanceName = c.Args().Get(0)
				action.Run(profile, instanceName)
			} else {
				fmt.Println("Instance is not selected")
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
