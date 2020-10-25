package main

import (
	"os"

	"github.com/diegomagdaleno/cheemit/cmd"
	log "github.com/sirupsen/logrus"
)

// Special thanks to the amazing logrus documentation
func init() {
	// Log as JSON instead of the default ASCII formatter.
	Formatter := new(log.TextFormatter)
	// Outoput to Stdout I would output to stderr but that isnt a very good idea
	log.SetOutput(os.Stdout)
	log.SetFormatter(Formatter)
}

func main() {
	cmd.Execute()
}
