package repository

import (
	"database/sql"
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

// получить пост по id
func (d *Database) GetPost(id int64) (*sql.Result, error) {
	post, err := d.db.NamedExec("select * from post where id = :id", id)
	if err != nil {
		return nil, fmt.Errorf("error getting post: %w", err)
	}
	return &post, nil
}

// удалить пост по id
func (d *Database) DeletePost(id int64) error {
	_, err := d.db.NamedExec("delete from post where id = :id", id)
	if err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}
	return nil
}
