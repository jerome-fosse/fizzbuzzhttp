package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/object-it/fizzbuzzhttp/fizzbuzzer"
	"github.com/sirupsen/logrus"
)

var version string = "undefined"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz", FizzbuzzHandler)

	srv := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	logrus.Fatal(srv.ListenAndServe())
}

// FizzbuzzHandler handles http requests to fizzbuzz
func FizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	fb := fizzbuzzer.New()

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, fb.Get())
}
