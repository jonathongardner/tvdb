package cli

import (
	"fmt"
	"os"

	"github.com/jonathongardner/dvddb/app"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func Run() error {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "logging level",
		},
	}

	app := &cli.App{
		// EnableBashCompletion: true,
		Name:    "dvddb",
		Version: app.Version,
		Usage:   "App for help naming ripped dvds!",
		Commands: []*cli.Command{
			tvCommand,
			mgCommand,
		},
		Flags: flags,
		Before: func(c *cli.Context) error {
			if c.Bool("verbose") {
				log.SetLevel(log.DebugLevel)
				log.Debug("Setting to debug...")
			}
			return nil
		},
	}
	return app.Run(os.Args)
}
