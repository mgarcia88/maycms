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

func (u *PostContentUseCase) Execute(content *entities.Content) (entities.Content, error) {
	cont, err := u.repo.CreateContent(content)
	if err != nil {
		return cont, err
	}
	return cont, nil
}
