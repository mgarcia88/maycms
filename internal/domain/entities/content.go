package entities

import (
	"errors"
	"time"
)

type Content struct {
	ID          int
	Title       string
	ContentText string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      string
}

func NewContent(t string, ct string, s string) (*Content, error) {
	if t == "" || len(t) < 10 {
		return &Content{}, errors.New("título inválido")
	}

	if ct == "" || len(ct) < 20 {
		return &Content{}, errors.New("conteúdo inválido")
	}

	if s == "" || len(s) < 5 {
		return &Content{}, errors.New("status inválido")
	}

	return &Content{Title: t, ContentText: ct, Status: s}, nil
}
