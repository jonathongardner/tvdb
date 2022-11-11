package cli

import (
	"fmt"
	"os"
	"github.com/jonathongardner/go-starter/app"

	"github.com/urfave/cli/v2"
)


func Run() (error) {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}

	app := &cli.App{
		Name: "starter",
		Version: app.Version,
		Usage: "Example starter app for cli tools!",
		Commands: []*cli.Command{
  		helloCommand,
  	},
	}
	return app.Run(os.Args)
}
