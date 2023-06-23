package models

import (
	"database/sql"
	"time"
)

type Post struct {
	Id          uint
	CategoryId  uint
	UserId      uint
	Title       string
	Slug        string
	Image       sql.NullString
	Excerpt     string
	Body        string
	PublishedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
