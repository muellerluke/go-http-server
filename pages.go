package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageSlug := vars["pageSlug"]
	log.Println(pageSlug)
	if pageSlug == "" {
		pageSlug = "index"
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, " + pageSlug))
}
