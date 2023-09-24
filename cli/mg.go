package cli

import (
	"fmt"
	"time"

	"github.com/jonathongardner/dvddb/routines"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

type greeting struct {
	name  string
	count int
}

func (g *greeting) Run(rc *routines.Controller) error {
	t := (2 * g.count) + 2
	time.Sleep(time.Duration(t) * time.Second)
	log.Infof("Hello %v (%v)", g.name, t)
	return nil
}

type waiting struct{}

func (w *waiting) Run(rc *routines.Controller) error {
	count := 0
	for {
		select {
		case <-rc.IsDone():
			return nil
		default:
			if count > 8 {
				return fmt.Errorf("To many people!")
			}
			time.Sleep(1 * time.Second)
			log.Info("Still waiting...")
			count++
		}
	}
	return nil
}

var mgCommand = &cli.Command{
	Name:      "many-greetings",
	Aliases:   []string{"m"},
	Usage:     "say many hellos",
	ArgsUsage: "[whos]",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "how-many",
			Value:   5,
			Usage:   "How many to say hello",
			EnvVars: []string{"START_HOW_MANY"},
			Action: func(ctx *cli.Context, v int) error {
				if 0 >= v {
					return fmt.Errorf("Flag number to say hello %v must be greater than 0", v)
				}
				return nil
			},
		},
	},
	Action: func(c *cli.Context) error {
		if c.NArg() == 0 {
			return fmt.Errorf("Must pass someone to talk to.")
		}

		routineController := routines.NewController()

		routineController.GoBackground(&waiting{})

		i := 0
		for i < c.NArg() {
			routineController.Go(&greeting{name: c.Args().Get(i), count: i})
			i++
		}

		return routineController.IsFinish()
	},
}
