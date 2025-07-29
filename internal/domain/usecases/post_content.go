package usecases

import (
	"maycms/internal/adapters/driven/infra/data/repositories"
	"maycms/internal/domain/entities"
)

type PostContentUseCase struct {
	repo repositories.ContentRepository
}

func NewPostContentUseCase(repo repositories.ContentRepository) *PostContentUseCase {
	return &PostContentUseCase{repo: repo}
}

func (u *PostContentUseCase) Execute(content *entities.Content) error {
	err := u.repo.CreateContent(content)
	if err != nil {
		return err
	}
	return nil
}
