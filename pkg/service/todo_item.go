package service

import (
	"REST_API"
	"REST_API/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userId, listId int, item REST_API.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	// list does not exist or don't belong to user
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, item)
}
