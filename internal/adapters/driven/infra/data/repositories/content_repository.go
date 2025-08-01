package repositories

import (
	entities "maycms/internal/domain/entities"
	ports "maycms/internal/domain/ports/driven"

	"github.com/sirupsen/logrus"
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
		logrus.WithFields(logrus.Fields{
			"contentId": id,
		}).Error("Failed to open database connection")
	}

	query := "SELECT c.id, c.title, c.content_text, c.status, c.created_at, c.updated_at, c.main_image, c.user_id, u.name, u.email FROM public.contents c INNER JOIN public.users u on u.id = c.user_id WHERE c.id = $1"

	defer c.db.CloseConnection(con)

	row := c.db.QueryRow(con, query, id)

	row.Scan(&content.ID, &content.Title, &content.ContentText, &content.Status, &content.CreatedAt, &content.UpdatedAt, &content.MainImage, &content.User.ID, &content.User.Name, &content.User.Email)

	return &content
}

func (c ContentRepository) GetAllContents() []entities.Content {
	var contents []entities.Content
	con, err := c.db.OpenConnection()

	if err != nil {
		logrus.WithField("error", err).Error("Failed to open database connection")
	}

	query := "SELECT c.id, c.title, c.content_text, c.status, c.created_at, c.updated_at, c.main_image, c.user_id, u.name, u.email FROM public.contents c INNER JOIN public.users u on u.id = c.user_id"

	defer c.db.CloseConnection(con)

	rows, err := c.db.Query(con, query)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"allContents": "all",
		}).Error("Failed to open database connection")
	}

	for rows.Next() {
		var content entities.Content

		err = rows.Scan(&content.ID, &content.Title, &content.ContentText, &content.Status, &content.CreatedAt, &content.UpdatedAt, &content.MainImage, &content.User.ID, &content.User.Name, &content.User.Email)

		if err != nil {
			logrus.WithFields(logrus.Fields{
				"contentList": content,
			}).Error("Failed listing contents")
			continue
		}

		contents = append(contents, content)
	}

	return contents
}

func (c ContentRepository) CreateContent(cont *entities.Content) error {

	con, err := c.db.OpenConnection()
	if err != nil {
		logrus.WithFields(logrus.Fields{"newContent": cont.Title}).Error("Failed to open database connection")
	}

	query := "INSERT INTO public.contents (title, content_text, status, user_id, main_image) VALUES($1, $2, $3, $4, $5);"

	defer c.db.CloseConnection(con)

	_, err = con.Exec(query, cont.Title, cont.ContentText, cont.Status, cont.User.ID, cont.MainImage)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"newContent": cont.Title,
		}).Error("Failed creating content")
		return err
	}

	return nil
}
