package service

import (
	"errors"
	todo "github.com/w-someproj/go-stuff/applications/todoapp"
	"github.com/w-someproj/go-stuff/applications/todoapp/pkg/repository"
)

// service for working with lists

type TodoListService struct {
	repos repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repos: repo}
}

func (s *TodoListService) CreateList(userId int, list todo.TodoList) (int, error) {
	return s.repos.CreateList(userId, list)
}

func (s *TodoListService) GetAllLists(userId int) ([]todo.TodoList, error) {
	return s.repos.GetAllLists(userId)
}

func (s *TodoListService) GetListById(userId int, listId int) (todo.TodoList, error) {
	return s.repos.GetListById(userId, listId)
}

func (s *TodoListService) DeleteListById(userId int, listId int) error {
	return s.repos.DeleteListById(userId, listId)

}

func (s *TodoListService) UpdateList(userId int, listid int, input todo.UpdateListInput) error {
	// if structure empty -> don`t update database
	if input.Title == nil && input.Description == nil {
		return errors.New(" no values for update")
	}
	return s.repos.UpdateList(userId, listid, input)
}
