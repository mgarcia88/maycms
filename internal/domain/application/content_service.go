package application

import (
	"maycms/internal/adapters/driven/infra/data/repositories"
	"maycms/internal/domain/entities"
)

type ContentService struct {
	repo repositories.ContentRepository
}

func NewContentService(repo repositories.ContentRepository) *ContentService {
	return &ContentService{repo: repo}
}

func (c *ContentService) GetContentById(id int) *entities.Content {
	content := c.repo.GetContentById(id)

	return content
}
