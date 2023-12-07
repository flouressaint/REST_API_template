package service

import (
	"REST_API"
	"REST_API/pkg/repository"
)

type Authorization interface {
	CreateUser(user REST_API.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list REST_API.TodoList) (int, error)
	GetAll(userId int) ([]REST_API.TodoList, error)
	GetById(userId, listId int) (REST_API.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input REST_API.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item REST_API.TodoItem) (int, error)
	GetAll(userId, listId int) ([]REST_API.TodoItem, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
