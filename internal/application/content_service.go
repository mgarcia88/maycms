package application

import (
	"maycms/internal/domain/entities"
	"maycms/internal/domain/usecases"
)

type ContentService struct {
	getAllContentsUseCase usecases.GetAllContentsUseCase
	getContentByIdUseCase usecases.GetContentByIdUseCase
	postContentUseCase    usecases.PostContentUseCase
}

func NewContentService(
	getAllContentsUseCase usecases.GetAllContentsUseCase,
	getContentByIdUsesCase usecases.GetContentByIdUseCase,
	postContentUseCase usecases.PostContentUseCase,
) *ContentService {
	return &ContentService{
		getAllContentsUseCase: getAllContentsUseCase,
		getContentByIdUseCase: getContentByIdUsesCase,
		postContentUseCase:    postContentUseCase,
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
	err := c.postContentUseCase.Execute(cont)

	return err
}
