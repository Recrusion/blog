package post

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

// создание нового поста
func NewPost(id int64, title string, content string, author string, tags []string) *Post {
	return &Post{
		ID:        id,
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Tags:      tags,
	}
}
