package repository

import (
	"REST_API"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user REST_API.User) (int, error)
	GetUser(username, password string) (REST_API.User, error)
}

type TodoList interface {
	Create(userId int, list REST_API.TodoList) (int, error)
	GetAll(userId int) ([]REST_API.TodoList, error)
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
		TodoList:      NewTodoListPostgres(db),
	}
}
