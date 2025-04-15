package dto

import (
	"time"
)

type CreatePostDto struct {
	ID        *string
	UserID    *string
	Content   *CreateContentDto
	Comments  []*CommentCreatePayload
	CreatedAt *time.Time
}
