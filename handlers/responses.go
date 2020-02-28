package handlers

import (
	"fmt"
	"net/http"
)

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
