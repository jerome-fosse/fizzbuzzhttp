package main

import (
	"fmt"
	"net/http"
)

func whoAmiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "FizzBuzz HTTP v%s written by Jérôme Fosse.", version)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 page not found")
	}
}
