package entities

import "time"

type Content struct {
	ID          int
	Title       string
	ContentText string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      string
}

func NewContent() *Content {
	return &Content{}
}
