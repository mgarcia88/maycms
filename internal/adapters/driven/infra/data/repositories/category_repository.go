package repositories

import (
	"maycms/internal/domain/entities"
	ports "maycms/internal/domain/ports"
)

type CategoryRepository struct {
	db ports.Database
}

func NewCategoryRepository(db ports.Database) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (c CategoryRepository) CreateCategory(cat entities.Category) error {
	con, err := c.db.OpenConnection()
	if err != nil {
		return err
	}

	query := "INSERT INTO public.categories (title, description) VALUES($1, $2);"

	defer c.db.CloseConnection(con)

	_, err = con.Exec(query, cat.Title, cat.Description)

	if err != nil {
		return err
	}

	return err
}
