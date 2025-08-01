package entities

import "errors"

type User struct {
	ID    int
	Name  string
	Email string
}

func NewUser(name string, email string) (*User, error) {
	if name == "" || len(name) < 3 {
		return &User{}, errors.New("invalid name")
	}

	if email == "" || len(email) < 5 {
		return &User{}, errors.New("invalid email")
	}

	return &User{Name: name, Email: email}, nil
}

func NewUserWithID(id int, name string, email string) User {
	user, _ := NewUser(name, email)

	user.ID = id
	return *user
}
