package service

import (
	todo "github.com/w-someproj/go-stuff/applications/todoapp"
	"github.com/w-someproj/go-stuff/applications/todoapp/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateList(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(userId int, listId int) (todo.TodoList, error)
	DeleteListById(userId int, listId int) error
	UpdateList(userId int, listid int, input todo.UpdateListInput) error
}

type TodoItem interface {
	CreateItem(userId int, listId int, item todo.TodoItem) (int, error)
	GetAllItems(userId int, listId int) ([]todo.TodoItem, error)
	GetItemById(userId int, itemId int) (todo.TodoItem, error)
	DeleteItemById(userId int, itemId int) error
	UpdateItem(userId int, itemId int, input todo.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// services need to call db
// depends from repos
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
