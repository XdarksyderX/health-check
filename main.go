package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "health-checker",
		Usage: "A tiny tool that checks wether a website is running or is down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name of the website to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Usage:       "Port number of the website to check",
				Required:    false,
				DefaultText: "80",
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("port") != "" {
				fmt.Println(Check(c.String("domain"), c.String("port")))
			} else {
				fmt.Println(Check(c.String("domain"), "80"))
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
