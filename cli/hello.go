package cli

import (
	"github.com/urfave/cli/v2"
	log "github.com/sirupsen/logrus"
)

var helloCommand =  &cli.Command{
	Name:    "hello",
	Aliases: []string{"b"},
	Usage:   "Say hello to someone",
	ArgsUsage: "[who]",
	Action:  func(c *cli.Context) error {
		who := c.Args().Get(0)
		if who == "" {
			who = "world"
		}
		log.Infof("Hello %v", who)

		return nil
	},
}
