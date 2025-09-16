package repositories

import (
	"fmt"
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

	query := "SELECT c.id, c.title, c.content_text, c.status, c.created_at, c.updated_at,  coalesce(c.main_image, 'https://i.scdn.co/image/ab67616d00001e02475ca6e5c1ce0ef70740c3c6') as main_image, c.user_id, u.name, u.email FROM public.contents c INNER JOIN public.users u on u.id = c.user_id WHERE c.id = $1"

	var content entities.Content

	err := c.db.GetDB().QueryRow(query, id).Scan(
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
	query := `
        SELECT c.id, c.title, c.content_text, c.status, c.created_at, c.updated_at, 
            	coalesce(c.main_image, 'https://i.scdn.co/image/ab67616d00001e02475ca6e5c1ce0ef70740c3c6') as main_image, 
				c.user_id, u.name, u.email
        FROM public.contents c
        INNER JOIN public.users u ON u.id = c.user_id
    `

	rows, err := c.db.GetDB().Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	defer func() {
		if cerr := rows.Close(); cerr != nil {
			logrus.WithError(cerr).Error("failed to close rows")
		}
	}()

	var contents []entities.Content

	for rows.Next() {
		var content entities.Content
		if err := rows.Scan(
			&content.ID,
			&content.Title,
			&content.ContentText,
			&content.Status,
			&content.CreatedAt,
			&content.UpdatedAt,
			&content.MainImage,
			&content.User.ID,
			&content.User.Name,
			&content.User.Email,
		); err != nil {
			logrus.WithError(err).Error("failed scanning content row")
			continue
		}

		contents = append(contents, content)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return contents, nil
}

func (c ContentRepository) CreateContent(cont *entities.Content) (entities.Content, error) {

	query := "INSERT INTO public.contents (title, content_text, status, user_id, main_image) VALUES($1, $2, $3, $4, $5) RETURNING id;"

	err := c.db.GetDB().QueryRow(query, cont.Title, cont.ContentText, cont.Status, cont.User.ID, cont.MainImage).Scan(&cont.ID)

	if err != nil {
		return *cont, err
	}

	return *cont, nil
}
