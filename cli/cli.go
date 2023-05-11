package cli

import (
	"fmt"
	"os"
	"github.com/jonathongardner/go-starter/app"

	"github.com/urfave/cli/v2"
	log "github.com/sirupsen/logrus"
)


func Run() (error) {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name: "version",
		Usage: "print the version",
	}

	flags := []cli.Flag {
		&cli.BoolFlag{
			Name: "verbose",
			Aliases: []string{"v"},
			Usage: "logging level",
		},
	}


	app := &cli.App{
		Name: "starter",
		Version: app.Version,
		Usage: "Example starter app for cli tools!",
		Commands: []*cli.Command{
			helloCommand,
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
