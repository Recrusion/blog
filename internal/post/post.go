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
func (p *Post) NewPost(ID int64, title string, content string, tags []string) *Post {
	return &Post{
		ID:        ID,
		Title:     title,
		Content:   content,
		Author:    p.Author,
		CreatedAt: time.Now(),
		Tags:      tags,
	}
}
