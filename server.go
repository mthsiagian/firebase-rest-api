package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, " up and running")
	})

	const port string = ":4000"
	log.Println("Server is started on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
