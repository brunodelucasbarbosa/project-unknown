package dto

import (
	"time"
)

type CommentCreatePayload struct {
	PostID    *string
	UserID    *string
	Content   *CreateContentDto
	CreatedAt *time.Time
}
