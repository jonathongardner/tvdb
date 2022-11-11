package cli

import (
	"github.com/urfave/cli/v2"
	log "github.com/sirupsen/logrus"
)

var helloCommand =  &cli.Command{
	Name:    "hello",
	Aliases: []string{"b"},
	Usage:   "backup photos in folder",
	Flags: []cli.Flag {
		&cli.StringFlag{
			Name:    "who",
			Aliases: []string{"w"},
			Value:   "world",
			EnvVars: []string{"USER"},
			Usage:   "Who to say hello to",
		},
	},
	Action:  func(c *cli.Context) error {
		who := c.String("who")
		log.Infof("Hello %v", who)

		return nil
	},
}
