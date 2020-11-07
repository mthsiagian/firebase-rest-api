package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/mthsiagian/firebase-rest-api/entity"
	"github.com/mthsiagian/firebase-rest-api/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error" : "Error geting posts"}`))
		return
	}

	if err = json.NewEncoder(res).Encode(posts); err != nil {
		http.Error(res, "Error encoding posts", http.StatusInternalServerError)
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}

func addPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		http.Error(res, "Error marshalling json", http.StatusInternalServerError)
		return
	}

	post.ID = rand.Int63()
	_, err = repo.Save(&post)

	if err != nil {
		http.Error(res, fmt.Sprintf("Error saving post: %v", err), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)
}
