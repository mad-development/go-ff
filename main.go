package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/feature", GetHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
