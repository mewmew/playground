package main

import (
	"flag"
	"log"
	"net/http"
)

func logger(handler http.Handler) http.Handler {
	logger := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(logger)
}

func main() {
	var addr string
	flag.StringVar(&addr, "http", ":8080", "HTTP service address (e.g., ':6060')")
	flag.Parse()
	handler := logger(http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatal(err)
	}
}
