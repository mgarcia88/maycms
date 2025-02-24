package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Content struct {
	ID          string    `valid:"uuid"`
	Title       string    `valid:"notnull"`
	ContentText string    `valid:"notnull"`
	Status      string    `valid:"notnull"`
	CreatedAt   time.Time `valid:"-"`
	UpdatedAt   time.Time `valid:"-"`
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
