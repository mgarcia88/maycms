package repositories

import "maycms/internal/domain/entities"

type ContentRepository interface {
	GetContentById() *entities.Content
	GetAllContents() *[]entities.Content
}
