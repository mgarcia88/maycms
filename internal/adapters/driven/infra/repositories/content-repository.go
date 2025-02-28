package repositories

import (
	"maycms/internal/domain/entities"
	"time"
)

type ContentRepository struct {
}

func NewContentRepository() *ContentRepository {
	return &ContentRepository{}
}

func (c ContentRepository) GetContentById(id int) entities.Content {
	content := *entities.NewContent()
	content.ContentText = "Lorem ipsum"
	content.Title = "É apenas um conteúdo"
	content.ID = id
	content.CreatedAt = time.Now()
	content.UpdatedAt = time.Now()
	content.Status = "Ativo"
	return content
}

func (c ContentRepository) GetAllContents() []entities.Content {
	mockContent := []entities.Content{
		1: {ID: 1, Title: "Meu primeiro conteudo", ContentText: "Lorem ipsum", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		2: {ID: 2, Title: "Meu segundo conteudo", ContentText: "Lorem ipsum", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	return mockContent
}
