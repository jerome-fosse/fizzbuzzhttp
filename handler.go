package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"

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
	int1, word1, int2, word2, limit := parseRequestParameters(r.URL)
	logrus.Infof("Requesting new Fizzbuzz with parameters int1=%d word1=%s int2=%d word2=%s limit=%d", int1, word1, int2, word2, limit)
	fb := fizzbuzzer.New(int1, int2, limit, word1, word2)

	body, err := json.Marshal(Fizzbuzz{fb.Get()})
	if err != nil {
		internalServerError(w, err)
	} else {
		okResponse(w, body)
	}
}

func parseRequestParameters(url *url.URL) (int1 int, word1 string, int2 int, word2 string, limit int) {
	int1, err := parseParamToInt(url, "int1")
	if err != nil {
		logrus.Debugf("int1: %s", err.Error())
		int1 = 3
	}

	int2, err = parseParamToInt(url, "int2")
	if err != nil {
		logrus.Debugf("int2: %s", err.Error())
		int2 = 5
	}

	limit, err = parseParamToInt(url, "limit")
	if err != nil {
		logrus.Debugf("limit: %s", err.Error())
		limit = 15
	}

	word1, err = parseParam(url, "word1")
	if err != nil {
		logrus.Debugf("word1: %s", err.Error())
		word1 = "fizz"
	}

	word2, err = parseParam(url, "word2")
	if err != nil {
		logrus.Debugf("word2: %s", err.Error())
		word2 = "buzz"
	}

	return
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
