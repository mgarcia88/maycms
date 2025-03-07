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

func (c *ContentService) GetAllContents() *[]entities.Content {
	contents := c.repo.GetAllContents()

	return contents
}

func (c *ContentService) CreateContent(cont *entities.Content) error {
	err := c.repo.CreateContent(cont)

	return err
}
