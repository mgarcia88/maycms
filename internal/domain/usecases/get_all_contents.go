package usecases

import (
	"maycms/internal/adapters/driven/infra/data/repositories"
	"maycms/internal/domain/entities"
)

type GetAllContentsUseCase struct {
	repo repositories.ContentRepository
}

func NewGetAllContentsUseCase(repo repositories.ContentRepository) *GetAllContentsUseCase {
	return &GetAllContentsUseCase{repo: repo}
}
func (u *GetAllContentsUseCase) Execute() ([]entities.Content, error) {
	contents := u.repo.GetAllContents()

	return contents, nil
}
