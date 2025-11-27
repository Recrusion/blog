package repository

import (
	"fmt"

	"github.com/Recrusion/blog-api/internal/post"
)

// создание поста в базе данных
func (d *Database) CreatePost(post *post.Post) error {
	_, err := d.db.NamedExec("insert into post (id, title, content, author, createdAt, updatedAt, tags) values (:id, :title, :content, :author, :createdAt, :updatedAt, :tags)", post)
	if err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}
