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
}

type TodoItem interface {
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
	}
}
