package main

import (
	"log"
	"net/http"

	"github.com/mewmew/playground/turtle"
)

func main() {
	err := tortoise()
	if err != nil {
		log.Fatalln(err)
	}
}

func tortoise() (err error) {
	// Initiate server.
	radicalSrv, err := turtle.NewRadicalServer()
	if err != nil {
		return err
	}
	http.HandleFunc("/css/", turtle.ServeData)
	http.HandleFunc("/images/", turtle.ServeData)
	http.Handle("/radicals/", radicalSrv)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}
