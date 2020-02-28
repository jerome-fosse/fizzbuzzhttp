package handlers

import (
	"fmt"
	"net/http"
)

// NewStatsHandler creates a handler that process stats requests
func NewStatsHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, "In construction")
	}
}
