package domain

import "time"

type Comment struct {
	ID        *string
	PostID    *string
	UserID    *string
	Content   *Content
	CreatedAt *time.Time
}
