package util

import (
	"github.com/urfave/cli"
)

type Cli struct {
	App        *cli.App
	AppName    string
	AppUsage   string
	AppVersion string
	AppFlags   []cli.Flag
	AppAction  func(c *cli.Context) error
}

func (c *Cli) Run(args []string) error {
	return c.App.Run(args)
}
