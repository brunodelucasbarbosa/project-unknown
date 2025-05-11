package dto

import "github.com/brunodelucasbarbosa/project-unknown/core/domain"

type CreateContentDto struct {
	PostID *string
	Image  *string
	Text   *string
}

func (p *CreateContentDto) ToDomain() *domain.Content {
	return &domain.Content{
		PostID: p.PostID,
		Image:  p.Image,
		Text:   p.Text,
	}
}
