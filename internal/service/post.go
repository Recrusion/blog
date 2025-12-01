package service

import (
	"errors"
	"time"

	"github.com/Recrusion/blog-api/internal/domain"
)

func (s *Service) CreatePost(post *domain.Post) error {
	if post.Title == "" || post.Content == "" || post.Author == "" || post.Tags == nil || len(post.Tags) == 0 {
		return errors.New("not all parameters are filled in")
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	err := s.service.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetPost(id int64) (*domain.Post, error) {
	if id == 0 {
		return nil, errors.New("id cannot be empty")
	}

	post, err := s.service.GetPost(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}
