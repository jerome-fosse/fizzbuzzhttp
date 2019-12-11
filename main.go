package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Soon Fizzbuzz here"))
}