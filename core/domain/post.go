package domain

import "time"

type Post struct {
	ID        *string
	UserID    *string
	Content   *Content
	Comments  []*Comment
	CreatedAt *time.Time
}
