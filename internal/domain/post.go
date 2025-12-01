package domain

import (
	"github.com/lib/pq"
	"time"
)

// структура поста
type Post struct {
	ID        int64          `json:"ID" db:"id"`
	Title     string         `json:"Title" db:"title"`
	Content   string         `json:"Content" db:"content"`
	Author    string         `json:"Author" db:"author"`
	CreatedAt time.Time      `json:"CreatedAt" db:"created_at"`
	UpdatedAt time.Time      `json:"UpdatedAt" db:"updated_at"`
	Tags      pq.StringArray `json:"Tags" db:"tags"`
}
