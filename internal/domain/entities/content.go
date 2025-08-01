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
	MainImage   string
	User        User
}

func NewContent(t string, ct string, s string, u User, im string) (*Content, error) {
	if t == "" || len(t) < 10 {
		return &Content{}, errors.New("título inválido")
	}

	if ct == "" || len(ct) < 20 {
		return &Content{}, errors.New("conteúdo inválido")
	}

	if s == "" || len(s) < 5 {
		return &Content{}, errors.New("status inválido")
	}

	if u.ID <= 0 {
		return &Content{}, errors.New("ID de usuário inválido")
	}

	if im == "" {
		im = "default.jpg" // Default image if none provided
	}

	return &Content{Title: t, ContentText: ct, Status: s, User: u, MainImage: im}, nil

}
