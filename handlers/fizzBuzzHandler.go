package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"

	"github.com/object-it/fizzbuzzhttp/fizzbuzzer"
	"github.com/object-it/fizzbuzzhttp/statistics"
)

// FizzbuzzResponse a struct
type FizzbuzzResponse struct {
	Values []string `json:"result"`
}

// FizzBuzzQueryParameters what else
type FizzBuzzQueryParameters struct {
	int1  int
	word1 string
	int2  int
	word2 string
	limit int
}

func (p FizzBuzzQueryParameters) String() string {
	return fmt.Sprintf("[int1 = %d, word1 = %s, int2 = %d, word2 = %s, limit = %d]", p.int1, p.word1, p.int2, p.word2, p.limit)
}

// NewFizzBuzzHandler creates a new handler that process fizzbuzz requests
func NewFizzBuzzHandler(statRepo *statistics.StatisticRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			logrus.Errorf("/fizzbuzz - method %v not allowed.", r.Method)
			methodNotAllowedResponse(w, r)
			return
		}

		params := parseQueryParameters(r.URL)
		logrus.Infof("Requesting new Fizzbuzz with parameters %v", params)

		go func() {
			statRepo.Save(params.String())
		}()

		fb := fizzbuzzer.New(params.int1, params.int2, params.limit, params.word1, params.word2)
		if body, err := json.Marshal(FizzbuzzResponse{fb.Get()}); err != nil {
			internalServerError(w, err)
		} else {
			okResponse(w, body)
		}
	}
}

func parseQueryParameters(url *url.URL) FizzBuzzQueryParameters {
	return FizzBuzzQueryParameters{
		parseParamToInt(url, "int1", 3),
		parseParamToString(url, "word1", "fizz"),
		parseParamToInt(url, "int2", 5),
		parseParamToString(url, "word2", "buzz"),
		parseParamToInt(url, "limit", 15),
	}
}
