package main

import (
	"github.com/sirupsen/logrus"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/object-it/fizzbuzzhttp/fizzbuzzer"
)

// Fizzbuzz a struct
type Fizzbuzz struct {
	Values []string `json:"result"`
}

// FizzbuzzHandler handles http requests to fizzbuzz
func FizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getFizzbuzz(w, r)
	} else {
		logrus.Errorf("/fizzbuzz - method %v not allowed.", r.Method)
		methodNotAllowedResponse(w, r)
	}
}

func getFizzbuzz(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Requesting new Fizzbuzz.")
	fb := fizzbuzzer.NewWithDefaultValues()

	body, err := json.Marshal(Fizzbuzz{fb.Get()})
	if err != nil {
		internalServerError(w, err)
	} else {
		okResponse(w, body)
	}
}

func okResponse(w http.ResponseWriter, body []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func internalServerError(w http.ResponseWriter, err error) {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, err.Error())
}

func methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "Method %v not allowed.", r.Method)
}
