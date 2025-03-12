package repositories

import (
	ports "maycms/Internal/Domain/Ports/Driven"
	"maycms/internal/domain/entities"
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
		panic("Não foi possível conectar")
	}

	query := "SELECT id, title, content_text, status FROM public.contents WHERE id = $1"

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

	query := "SELECT id, title, content_text, status FROM public.contents"

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

func (c ContentRepository) CreateContent(cont *entities.Content) error {

	con, err := c.db.OpenConnection()
	if err != nil {
		panic("Não foi possível conectar")
	}

	query := "INSERT INTO public.contents (title, content_text, status) VALUES($1, $2, $3);"

	_, err = con.Exec(query, cont.Title, cont.ContentText, cont.Status)

	c.db.CloseConnection(con)

	if err != nil {
		return err
	}

	return nil
}
