package dto

import (
	"time"

	"github.com/brunodelucasbarbosa/project-unknown/core/domain"
)

type CommentCreatePayload struct {
	PostID    *string
	UserID    *string
	Content   *CreateContentDto
	CreatedAt *time.Time
}

func (c *CommentCreatePayload) ToDomain() *domain.Comment {
	return &domain.Comment{
		PostID:    c.PostID,
		UserID:    c.UserID,
		Content:   c.Content.ToDomain(),
		CreatedAt: c.CreatedAt,
	}
}
