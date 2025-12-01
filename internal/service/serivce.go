package service

import (
	"time"

	"github.com/Recrusion/blog-api/internal/domain"
	"github.com/Recrusion/blog-api/internal/repository"
)

type Service struct {
	service *repository.Database
}

func NewService(db *repository.Database) *Service {
	return &Service{
		service: db,
	}
}

// интерфейс слоя взаимодействия с базой данных
type PostService interface {
	CreatePost(post *domain.Post) error
	GetPost(id int64) (*domain.Post, error)
	DeletePost(id int64) error
	UpdatePost(id int64, title, content string, updatedAt time.Time, tags []string) error
}
