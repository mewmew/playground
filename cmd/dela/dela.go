package main

import (
	"log"
	"net/http"
)

func Log(handler http.Handler) http.Handler {
	logger := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(logger)
}

func main() {
	if err := http.ListenAndServe(":8080", Log(http.FileServer(http.Dir(".")))); err != nil {
		log.Fatal(err)
	}
}
