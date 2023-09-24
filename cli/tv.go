package cli

import (
	"fmt"
	"os"

	"github.com/jonathongardner/dvddb/tvdb"

	"github.com/urfave/cli/v2"
)

var tvCommand = &cli.Command{
	Name:      "tv",
	Usage:     "Name TV mkv file for plex",
	ArgsUsage: "[folder]",
	Subcommands: []*cli.Command{
		renameCommand,
		showDBCommand,
	},
}

var renameCommand = &cli.Command{
	Name:      "rename",
	Usage:     "Name TV mkv file for plex",
	ArgsUsage: "[folder]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "output-dir",
			Value:   "",
			Usage:   "Directory to move movies to",
			EnvVars: []string{"TV_OUTPUT_DIR"},
		},
	},
	Action: func(c *cli.Context) error {
		path := c.Args().Get(0)
		if path == "" {
			return fmt.Errorf("Path argument required")
		}

		destDir := c.String("output-dir")
		if destDir == "" {
			path, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("Error getting local path %v", err)
			}
			destDir = path
		}

		return tvdb.MoveFiles(destDir, path)
	},
}

var showDBCommand = &cli.Command{
	Name:  "show-db",
	Usage: "Show TV DB",
	Action: func(c *cli.Context) error {
		return tvdb.PrintDB()
	},
}
