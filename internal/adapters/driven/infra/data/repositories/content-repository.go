package repositories

import (
	data "maycms/internal/adapters/driven/infra/data/interfaces"
	"maycms/internal/domain/entities"
	"time"
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

func (c ContentRepository) GetAllContents() []entities.Content {
	mockContent := []entities.Content{
		1: {ID: 1, Title: "Meu primeiro conteudo", ContentText: "Lorem ipsum", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		2: {ID: 2, Title: "Meu segundo conteudo", ContentText: "Lorem ipsum", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	return mockContent
}
