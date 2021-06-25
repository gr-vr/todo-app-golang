package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/satanCoffee/todo-app-golang"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(email, password string) (todo.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
