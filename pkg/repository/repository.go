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
	GetById(userId, listId int) (REST_API.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input REST_API.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item REST_API.TodoItem) (int, error)
	GetAll(listId int) ([]REST_API.TodoItem, error)
	GetById(userId, itemId int) (REST_API.TodoItem, error)
	Update(userId, itemId int, input REST_API.UpdateItemInput) error
	Delete(userId, itemId int) error
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
		TodoItem:      NewTodoItemPostgres(db),
	}
}
