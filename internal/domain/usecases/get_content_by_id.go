package usecases

import (
	"maycms/internal/adapters/driven/infra/data/repositories"
	"maycms/internal/domain/entities"
)

type GetContentByIdUseCase struct {
	repo repositories.ContentRepository
}

func NewGetContentByIdUseCase(repo repositories.ContentRepository) *GetContentByIdUseCase {
	return &GetContentByIdUseCase{repo: repo}
}

func (u *GetContentByIdUseCase) Execute(id int) (*entities.Content, error) {
	content := u.repo.GetContentById(id)

	return content, nil
}
