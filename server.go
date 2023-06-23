package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/file-upload", VideoUploadHandler).Methods("POST")
	router.HandleFunc("/api/{apiSlug}", ApiHandler)
	router.HandleFunc("/{pageSlug}", PageHandler).Methods("GET")
	http.ListenAndServe(":80", router)
	log.Println("Listening on port 80")
}
