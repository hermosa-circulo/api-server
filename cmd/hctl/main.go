package main

import (
	"os"

	hctl "github.com/hermosa-circulo/iga-controller/cmd/hctl/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Initialize README.md and manage templates"
	app.EnableBashCompletion = true
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "get object",
			Subcommands: []cli.Command{
				{
					Name:   "gene",
					Usage:  "get one gene",
					Action: hctl.GetGene,
				},
				{
					Name:   "evaluation",
					Usage:  "get evaluation",
					Action: hctl.GetEvaluation,
				},
			},
		},
		{
			Name:  "set",
			Usage: "set object",
			Subcommands: []cli.Command{
				{
					Name:   "evaluation",
					Usage:  "set evaluation",
					Action: hctl.SetEvaluation,
				},
			},
		},
	}
	app.Run(os.Args)
}
