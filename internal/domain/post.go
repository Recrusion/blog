package domain

import "time"

// структура поста
type Post struct {
	ID        int64     `json:"ID"`
	Title     string    `json:"Title"`
	Content   string    `json:"Content"`
	Author    string    `json:"Author"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	Tags      []string  `json:"Tags"`
}

// интерфейс слоя взаимодействия с базой данных
type PostRepository interface {
	CreatePost(post *Post) error
	GetPost(id int64) (*Post, error)
	DeletePost(id int64) error
	UpdatePost(id int64, title, content string, updatedAt time.Time, tags []string) error
}
