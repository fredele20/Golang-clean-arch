package controller

import (
	"Demo/entity"
	"Demo/errors"
	"Demo/service/posts"
	"encoding/json"
	"net/http"
)

type PostController interface {
	AddPost(w http.ResponseWriter, r *http.Request)
	GetPosts(w http.ResponseWriter, r *http.Request)
}

type postController struct {
	postService posts.PostService
}

func NewPostService(s posts.PostService) PostController {
	return &postController{
		postService: s,
	}
}

func (c *postController) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.MessageError{Message:"error unmarshalling data"})
		return
	}
	result, err := c.postService.Create(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.MessageError{Message:err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}

func (c *postController) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []entity.Post
	result, err := c.postService.GetPosts(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.MessageError{Message:"error getting the posts"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
