package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var version = "undefined"
var build = "undefined"
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

	logrus.Infof("Starting Fizzbuzz v%s Build: %s", version, build)
	logrus.Debugf("Port=%d, Verbose=%v", port, verbose)

	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz", FizzbuzzHandler)

	srv := http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Infof("Fizzbuzz is started and listening on port %d.", port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	logrus.Info("Stopping Fizzbuzz.")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error while shutting down Fizzbuzz. %v\n", err)
	}

	logrus.Infoln("Fizzbuzz is stopped.")
}
