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

func (s *TodoItemService) GetAll(userId, listId int) ([]REST_API.TodoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	// list does not exist or don't belong to user
	if err != nil {
		return nil, err
	}
	return s.repo.GetAll(listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (REST_API.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input REST_API.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
