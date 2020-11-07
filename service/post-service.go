package service

import (
	"math/rand"

	"github.com/mthsiagian/firebase-rest-api/entity"
	"github.com/mthsiagian/firebase-rest-api/repository"
)

// PostService interface
type PostService interface {
	AddPost(post *entity.Post) error
	FindAll() ([]entity.Post, error)
}

type service struct {
	Post repository.PostRepository
}

// NewPostService to create instance of PostService
func NewPostService(post repository.PostRepository) PostService {
	return &service{
		Post: post,
	}
}

func (s *service) AddPost(post *entity.Post) error {
	post.ID = rand.Int63()
	_, err := s.Post.Save(post)
	return err

}

func (s *service) FindAll() ([]entity.Post, error) {
	return s.Post.FindAll()
}
