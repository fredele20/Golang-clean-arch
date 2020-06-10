package posts

import (
	"Demo/entity"
	"Demo/repository"
)

type PostService interface {
	Create(post *entity.Post) (*entity.Post, error)
	GetPosts(posts []entity.Post) ([]entity.Post, error)
	GetPost(post *entity.Post) (*entity.Post, error)
}

type postService struct{
	serviceRepo repository.PostRepository
}

func NewPostService(r repository.PostRepository) PostService {
	return &postService{
		serviceRepo: r,
	}
}

func (s *postService) Create(post *entity.Post) (*entity.Post, error) {
	return s.serviceRepo.Create(post)
}

func (s *postService) GetPosts(posts []entity.Post) ([]entity.Post, error) {
	return s.serviceRepo.FindAllPost(posts)
}

func (s *postService) GetPost(post *entity.Post) (*entity.Post, error) {
	return s.serviceRepo.FindOnePost(post)
}


