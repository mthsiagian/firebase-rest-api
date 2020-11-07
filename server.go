package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		fmt.Println(e)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, " up and runningg")
	})

	router.HandleFunc("/post", getPosts).Methods("GET")
	router.HandleFunc("/post", addPost).Methods("POST")

	const port string = ":4000"
	log.Println("Server is started on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
