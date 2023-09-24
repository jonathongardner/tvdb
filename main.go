package main

import (
	"github.com/jonathongardner/dvddb/cli"

	log "github.com/sirupsen/logrus"
)

func main() {
	err := cli.Run()
	if err != nil {
		log.Fatal(err)
	}
}
