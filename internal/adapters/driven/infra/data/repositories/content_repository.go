package repositories

import (
	"log"
	entities "maycms/internal/domain/entities"
	ports "maycms/internal/domain/ports/driven"
)

type ContentRepository struct {
	db ports.Database
}

func NewContentRepository(db ports.Database) *ContentRepository {
	return &ContentRepository{db: db}
}

func (c ContentRepository) GetContentById(id int) *entities.Content {
	var content entities.Content
	con, err := c.db.OpenConnection()

	if err != nil {
		log.Fatal("Não foi possível conectar")
	}

	query := "SELECT id, title, content_text, status, created_at, updated_at FROM public.contents WHERE id = $1"

	defer c.db.CloseConnection(con)

	row := c.db.QueryRow(con, query, id)

	row.Scan(&content.ID, &content.Title, &content.ContentText, &content.Status, &content.CreatedAt, &content.UpdatedAt)

	return &content
}

func (c ContentRepository) GetAllContents() *[]entities.Content {
	var contents []entities.Content
	con, err := c.db.OpenConnection()

	if err != nil {
		log.Fatal("Não foi possível conectar")
	}

	query := "SELECT id, title, content_text, status, created_at, updated_at FROM public.contents"

	defer c.db.CloseConnection(con)

	rows, err := c.db.Query(con, query)

	if err != nil {
		log.Fatal("Não foi possível conectar")
	}

	for rows.Next() {
		var content entities.Content

		err = rows.Scan(&content.ID, &content.Title, &content.ContentText, &content.Status, &content.CreatedAt, &content.UpdatedAt)

		if err != nil {
			continue
		}

		contents = append(contents, content)
	}

	return &contents
}

func (c ContentRepository) CreateContent(cont *entities.Content) error {

	con, err := c.db.OpenConnection()
	if err != nil {
		log.Fatal("Não foi possível conectar")
	}

	query := "INSERT INTO public.contents (title, content_text, status) VALUES($1, $2, $3);"

	defer c.db.CloseConnection(con)

	_, err = con.Exec(query, cont.Title, cont.ContentText, cont.Status)

	if err != nil {
		return err
	}

	return nil
}
