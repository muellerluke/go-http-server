package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apiSlug := vars["apiSlug"]
	log.Println(apiSlug)
	if apiSlug == "" {
		apiSlug = "index"
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, " + apiSlug))
}
