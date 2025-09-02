package ports

import "maycms/internal/domain/entities"

type ContentRepository interface {
	GetContentById(id int) (*entities.Content, error)
	GetAllContents() ([]entities.Content, error)
	CreateContent(content *entities.Content) (entities.Content, error)
}
