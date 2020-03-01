package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/object-it/fizzbuzzhttp/handlers"
	"github.com/object-it/fizzbuzzhttp/statistics"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var version = "undefined"
var build = "undefined"

// FizzBuzzServer a FizzBuzz Server
type FizzBuzzServer struct {
	port       int
	httpServer http.Server
}

// NewFizzBuzzServer creates a new Server
func NewFizzBuzzServer(port int) *FizzBuzzServer {
	repo := statistics.NewRepository()
	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz", handlers.NewFizzBuzzHandler(repo))
	router.HandleFunc("/whoami", handlers.NewWhoAmIHandler(version))
	router.HandleFunc("/stats", handlers.NewStatsHandler(repo))

	return &FizzBuzzServer{
		port,
		http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: router,
		},
	}
}

// Start a FizzBuzz Server
func (s *FizzBuzzServer) Start() {
	logrus.Infof("Starting Fizzbuzz v%s Build: %s", version, build)

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Infof("Fizzbuzz is started and listening on port %d.", s.port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	logrus.Info("Stopping Fizzbuzz.")
	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error while shutting down Fizzbuzz. %v\n", err)
	}

	logrus.Infoln("Fizzbuzz is stopped.")
}
