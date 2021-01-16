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
				return glitzy.Add()
			},
		},
		{
			Name:        "wipe",
			Usage:       "wipe the entire passwords",
			Description: "wipe the entire passwords",
			Action: func(c *cli.Context) error {
				return glitzy.Wipe()
			},
		},
		{
			Name:        "search",
			Usage:       "Search for password",
			Description: "Search for password",
			Action: func(c *cli.Context) error {
				return glitzy.Search()
			},
		},
	}
	return app.Run(os.Args)
}
