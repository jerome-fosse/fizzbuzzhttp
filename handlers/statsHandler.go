package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/object-it/fizzbuzzhttp/statistics"
	"github.com/sirupsen/logrus"
)

const (
	uintSize = 32 << (^uint(0) >> 32 & 1)
	maxInt   = 1<<(uintSize-1) - 1
)

// NewStatsHandler creates a handler that process stats requests
func NewStatsHandler(statRepo *statistics.StatisticRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			logrus.Errorf("/stats - method %v not allowed.", r.Method)
			methodNotAllowedResponse(w, r)
			return
		}

		limit := parseParamToInt(r.URL, "limit", maxInt)
		stats := statRepo.FindTopLimitBy(limit)
		if body, err := json.Marshal(stats); err != nil {
			internalServerError(w, err)
		} else {
			okResponse(w, body)
		}
	}
}
