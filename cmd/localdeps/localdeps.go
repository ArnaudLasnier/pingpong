package main

import (
	"os"

	"github.com/ArnaudLasnier/pingpong/internal/tools/localdeps"
	"github.com/urfave/cli/v2"
)

func main() {
	var err error
	cmd := &cli.App{
		Name: "localdeps",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "env-file",
				Required: true,
			},
		},
		Action: localdeps.Run,
	}
	err = cmd.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
