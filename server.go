package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mthsiagian/firebase-rest-api/service"

	"github.com/joho/godotenv"
	"github.com/mthsiagian/firebase-rest-api/controller"
	"github.com/mthsiagian/firebase-rest-api/repository"

	"github.com/gorilla/mux"
)

var (
	pr repository.PostRepository = repository.NewPostRepository()
	ps service.PostService       = service.NewPostService(pr)
	pc controller.PostController = controller.NewPostController(ps)
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

	router.HandleFunc("/post", pc.GetPosts).Methods("GET")
	router.HandleFunc("/post", pc.AddPost).Methods("POST")

	const port string = ":4000"
	log.Println("Server is started on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
