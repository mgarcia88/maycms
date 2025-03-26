package entities

import (
	"errors"
	"time"
)

type Category struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCategorie(t string, d string) (*Category, error) {
	if t == "" || len(t) < 10 {
		return &Category{}, errors.New("título inválido")
	}

	if d == "" || len(d) < 20 {
		return &Category{}, errors.New("descrição inválida")
	}

	return &Category{Title: t, Description: d}, nil
}
