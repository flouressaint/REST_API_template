package repository

import (
	"REST_API"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user REST_API.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
