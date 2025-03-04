package repositories

import (
	data "maycms/internal/adapters/driven/infra/data/interfaces"
	"maycms/internal/domain/entities"
)

type ContentRepository struct {
	db data.Database
}

func NewContentRepository(db data.Database) *ContentRepository {
	return &ContentRepository{db: db}
}

func (c ContentRepository) GetContentById(id int) *entities.Content {
	var content entities.Content
	con, err := c.db.OpenConnection()

	if err != nil {
		panic("Não foi possível conectar")
	}

	query := "SELECT id, title, content_text, status FROM public.content WHERE id = $1"

	row := c.db.QueryRow(con, query, id)

	row.Scan(&content.ID, &content.Title, &content.ContentText, &content.Status)

	c.db.CloseConnection(con)

	return &content
}

func (c ContentRepository) GetAllContents() *[]entities.Content {
	var contents []entities.Content
	con, err := c.db.OpenConnection()

	if err != nil {
		panic("Não foi possível conectar")
	}

	query := "SELECT id, title, content_text, status FROM public.content"

	rows, err := c.db.Query(con, query)

	if err != nil {
		panic("Erro ao buscar os conteúdos")
	}

	for rows.Next() {
		var content entities.Content

		err = rows.Scan(&content.ID, &content.Title, &content.ContentText, &content.Status)

		if err != nil {
			continue
		}

		contents = append(contents, content)
	}

	c.db.CloseConnection(con)

	return &contents
}
