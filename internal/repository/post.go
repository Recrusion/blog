package repository

import (
	"fmt"
	"time"

	"github.com/Recrusion/blog-api/internal/domain"
)

// создание поста в базе данных
func (d *Database) CreatePost(post *domain.Post) error {
	_, err := d.db.NamedExec("insert into post (title, content, author, created_at, updated_at, tags) values (:title, :content, :author, :created_at, :updated_at, :tags) returning id", post)
	if err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}

// получить пост по id
func (d *Database) GetPost(id int64) (*domain.Post, error) {
	var post domain.Post
	err := d.db.Get(&post, "select * from post where id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("error getting post: %w", err)
	}
	return &post, nil
}

// удалить пост по id
func (d *Database) DeletePost(id int64) error {
	_, err := d.db.Exec("delete from post where id = :id", id)
	if err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}
	return nil
}

// обновить пост
func (d *Database) UpdatePost(id int64, title, content string, updatedAt time.Time, tags []string) error {
	_, err := d.db.NamedExec("update post set title = :title, content = :content, updated_at = :updatedAt, tags = :tags where id = :id",
		map[string]interface{}{
			"title":      title,
			"content":    content,
			"updated_at": updatedAt,
			"tags":       tags,
			"id":         id})

	if err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}
	return nil
}
