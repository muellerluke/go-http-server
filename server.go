package server

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	fmt.Println("Calling cmd.Execute()!")
}
