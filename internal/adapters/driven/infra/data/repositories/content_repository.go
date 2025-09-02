package repositories

import (
	entities "maycms/internal/domain/entities"
	ports "maycms/internal/domain/ports"

	"github.com/sirupsen/logrus"
)

type ContentRepository struct {
	db ports.Database
}

func NewContentRepository(db ports.Database) ContentRepository {
	return ContentRepository{db: db}
}

func (c ContentRepository) GetContentById(id int) (*entities.Content, error) {

	con, err := c.db.OpenConnection()

	if err != nil {
		return nil, err
	}

	defer c.db.CloseConnection(con)

	query := "SELECT c.id, c.title, c.content_text, c.status, c.created_at, c.updated_at, c.main_image, c.user_id, u.name, u.email FROM public.contents c INNER JOIN public.users u on u.id = c.user_id WHERE c.id = $1"

	var content entities.Content

	err = c.db.QueryRow(con, query, id).Scan(
		&content.ID,
		&content.Title,
		&content.ContentText,
		&content.Status,
		&content.CreatedAt,
		&content.UpdatedAt,
		&content.MainImage,
		&content.User.ID,
		&content.User.Name,
		&content.User.Email)

	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (c ContentRepository) GetAllContents() ([]entities.Content, error) {
	var contents []entities.Content
	con, err := c.db.OpenConnection()

	if err != nil {
		return nil, err
	}

	query := "SELECT c.id, c.title, c.content_text, c.status, c.created_at, c.updated_at, c.main_image, c.user_id, u.name, u.email FROM public.contents c INNER JOIN public.users u on u.id = c.user_id"

	defer c.db.CloseConnection(con)

	rows, err := c.db.Query(con, query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var content entities.Content

		err = rows.Scan(&content.ID, &content.Title, &content.ContentText, &content.Status, &content.CreatedAt, &content.UpdatedAt, &content.MainImage, &content.User.ID, &content.User.Name, &content.User.Email)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"contentList": content,
			}).Error(err, "Failed scanning content row")
			continue
		}

		contents = append(contents, content)
	}

	return contents, nil
}

func (c ContentRepository) CreateContent(cont *entities.Content) (entities.Content, error) {

	con, err := c.db.OpenConnection()
	if err != nil {
		return *cont, err
	}

	query := "INSERT INTO public.contents (title, content_text, status, user_id, main_image) VALUES($1, $2, $3, $4, $5) RETURNING id;"

	defer c.db.CloseConnection(con)

	err = con.QueryRow(query, cont.Title, cont.ContentText, cont.Status, cont.User.ID, cont.MainImage).Scan(&cont.ID)

	if err != nil {
		return *cont, err
	}

	return *cont, nil
}
