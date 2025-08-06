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
	contents, err := u.repo.GetAllContents()

	if err != nil {
		return nil, err
	}

	return contents, nil
}
