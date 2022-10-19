package service

import (
	"errors"
	todo "github.com/w-someproj/go-stuff/applications/todoapp"
	"github.com/w-someproj/go-stuff/applications/todoapp/pkg/repository"
)

type TodoItemService struct {
	repos     repository.TodoItem
	reposList repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, reposList repository.TodoList) *TodoItemService {
	return &TodoItemService{repos: repo, reposList: reposList}
}

func (s *TodoItemService) CreateItem(userId int, listId int, item todo.TodoItem) (int, error) {
	// if list don`t exist or don`t belong to user
	_, err := s.reposList.GetListById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repos.CreateItem(listId, item)
}

func (s *TodoItemService) GetAllItems(userId int, listId int) ([]todo.TodoItem, error) {
	// don`t check list and user now so can write bigger db query
	// but in real project - need check as in CreateItem
	return s.repos.GetAllItems(userId, listId)
}

func (s *TodoItemService) GetItemById(userId int, itemId int) (todo.TodoItem, error) {
	// get todoitem for user (not interested in list)
	return s.repos.GetItemById(userId, itemId)
}

func (s *TodoItemService) DeleteItemById(userId int, itemId int) error {
	return s.repos.DeleteItemById(userId, itemId)
}

func (s *TodoItemService) UpdateItem(userId int, itemId int, input todo.UpdateItemInput) error {
	// if structure empty -> don`t update database
	if input.Title == nil && input.Description == nil && input.Done == nil && input.Importance == nil {
		return errors.New(" no values for update")
	}
	return s.repos.UpdateItem(userId, itemId, input)
}
