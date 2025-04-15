package service

import "github.com/brunodelucasbarbosa/project-unknown/core/domain"

type PostService struct{}

func (p *PostService) CreatePost(userID string, post *domain.Post) error {
	// Implement the logic to create a post for a user
	return nil
}