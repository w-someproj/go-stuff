package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/w-someproj/go-stuff/applications/todoapp"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string, password string) (todo.User, error)
}

type TodoList interface {
	CreateList(userId int, list todo.TodoList) (int, error)
	GetAllLists(userId int) ([]todo.TodoList, error)
	GetListById(userId int, listId int) (todo.TodoList, error)
	DeleteListById(userId int, listId int) error
	UpdateList(userId int, listid int, input todo.UpdateListInput) error
}

type TodoItem interface {
	CreateItem(listId int, item todo.TodoItem) (int, error)
	GetAllItems(userId int, listId int) ([]todo.TodoItem, error)
	GetItemById(userId int, itemId int) (todo.TodoItem, error)
	DeleteItemById(userId int, itemId int) error
	UpdateItem(userId int, itemId int, input todo.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// initialize repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
