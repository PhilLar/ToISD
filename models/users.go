package models

import (
	sq "github.com/Masterminds/squirrel"
)

type User struct {
	Name     string
	Email    string
	Address  string
	Password string
	Phone    string
}

func (s *Store) InsertUser(user User) error {
	_, err := sq.Insert("users").
		Columns("name", "email", "address", "password", "phone_number").
		Values(user.Name, user.Email, user.Address, user.Password, user.Phone).
		RunWith(s.DB).
		PlaceholderFormat(sq.Dollar).Query()
	if err != nil {
		return err
	}
	return nil
}
