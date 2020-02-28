package main

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var port int
var verbose bool

func init() {
	flag.IntVar(&port, "port", 8080, "http port fizzbuzz is listening to")
	flag.BoolVar(&verbose, "verbose", false, "log debug informations")

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	flag.Parse()

	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.Debugf("Configuration : Port=%d, Verbose=%v", port, verbose)

	srv := NewFizzBuzzServer(port)
	srv.Start()
}
