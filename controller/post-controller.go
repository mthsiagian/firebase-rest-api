package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mthsiagian/firebase-rest-api/entity"
	"github.com/mthsiagian/firebase-rest-api/service"
)

// PostController interface
type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	postService service.PostService
}

// NewPostController to create PostController instance
func NewPostController(ps service.PostService) PostController {
	return &controller{
		postService: ps,
	}
}

func (c *controller) GetPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	posts, err := c.postService.FindAll()
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

func (c *controller) AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		http.Error(res, "Error marshalling json", http.StatusInternalServerError)
		return
	}

	err = c.postService.AddPost(&post)

	if err != nil {
		http.Error(res, fmt.Sprintf("Error saving post: %v", err), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)
}
