package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Content struct {
	ID          string
	Title       string
	ContentText string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewContent() *Content {
	return &Content{}
}

func (content *Content) Validate() error {
	_, err := govalidator.ValidateStruct(content)

	if err != nil {
		return err
	}

	return nil
}
