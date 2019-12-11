package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var version = "undefined"
var build = "undefined"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	logrus.Infof("Starting Fizzbuzz v%s Build: %s", version, build)

	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz", FizzbuzzHandler)

	srv := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Info("Fizzbuzz is started and ready to serve requests.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	logrus.Info("Stopping Fizzbuzz.")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error while shutting down Fizzbuzz. %v\n", err)
	}

	logrus.Infoln("Fizzbuzz is stopped.")
}
