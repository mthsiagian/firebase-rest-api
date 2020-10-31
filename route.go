package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{ID: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	result, err := json.Marshal(posts)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error" : "Error parsing the posts array"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func addPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "error" : "Error parsing post"}`))
		return
	}

	post.ID = len(posts) + 1
	posts = append(posts, post)
	fmt.Println(posts)
	res.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	res.Write(result)
}
