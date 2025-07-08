package application

import (
	"maycms/internal/adapters/driven/infra/data/repositories"
	"maycms/internal/domain/entities"
	"maycms/internal/domain/usecases"
)

type ContentService struct {
	repo                  repositories.ContentRepository
	getAllContentsUseCase usecases.GetAllContentsUseCase
	getContentByIdUseCase usecases.GetContentByIdUseCase
}

func NewContentService(repo repositories.ContentRepository,
	getAllContentsUseCase usecases.GetAllContentsUseCase,
	getContentByIdUsesCase usecases.GetContentByIdUseCase) *ContentService {
	return &ContentService{repo: repo,
		getAllContentsUseCase: getAllContentsUseCase,
		getContentByIdUseCase: getContentByIdUsesCase,
	}
}

func (c *ContentService) GetContentById(id int) *entities.Content {
	content, _ := c.getContentByIdUseCase.Execute(id)

	return content
}

func (c *ContentService) GetAllContents() *[]entities.Content {
	contents, _ := c.getAllContentsUseCase.Execute()

	return &contents
}

func (c *ContentService) CreateContent(cont *entities.Content) error {
	err := c.repo.CreateContent(cont)

	return err
}
