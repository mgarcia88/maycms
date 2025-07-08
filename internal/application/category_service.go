package application

import (
	repositories "maycms/internal/adapters/driven/infra/data/repositories"
	entities "maycms/internal/domain/entities"
)

type CategoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (c CategoryService) CreateCategory(cat entities.Category) error {
	err := c.repo.CreateCategory(cat)

	return err
}
