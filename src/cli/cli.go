package cli

import (
	"os"
	"time"

	"github.com/iamstefin/glitzy/src/glitzy"
	"github.com/urfave/cli/v2"
)

// Run will run the command line program
func Run() (err error) {
	app := cli.NewApp()
	app.Name = "glitzy"
	app.Compiled = time.Now()
	app.Usage = "a simple password manager"
	app.Commands = []*cli.Command{
		{
			Name:        "add",
			Usage:       "add a new password",
			Description: "add a new password",
			Action: func(c *cli.Context) error {
				glitzy.Add()
				return nil
			},
		},
		{
			Name:        "wipe",
			Usage:       "wipe the entire passwords",
			Description: "wipe the entire passwords",
			Action: func(c *cli.Context) error {
				glitzy.Wipe()
				return nil
			},
		},
	}
	return app.Run(os.Args)
}
