package dto

import (
	"time"

	"github.com/brunodelucasbarbosa/project-unknown/core/domain"
)

type CreatePostDto struct {
	UserID  *string
	Content *CreateContentDto
}

func (p *CreatePostDto) ToDomain() *domain.Post {
	n := time.Now()
	return &domain.Post{
		ID:        p.UserID,
		UserID:    p.UserID,
		Content:   p.Content.ToDomain(),
		Comments:  nil,
		CreatedAt: &n,
	}
}
