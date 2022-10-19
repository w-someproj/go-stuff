package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "github.com/w-someproj/go-stuff/applications/todoapp"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) CreateItem(listId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf(`INSERT INTO %s (title, description, importance) values ($1, $2, $3) RETURNING id`, todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description, item.Importance)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf(`INSERT INTO %s (list_id, item_id) values ($1, $2)`, listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAllItems(userId int, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done, ti.importance  FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id 
								WHERE ul.user_id = $1 AND li.list_id = $2 `, todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, userId, listId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetItemById(userId int, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done, ti.importance  FROM %s ti INNER JOIN %s li ON li.item_id = ti.id INNER JOIN %s ul ON ul.list_id = li.list_id 
								WHERE ul.user_id = $1 AND ti.id = $2 `, todoItemsTable, listsItemsTable, usersListsTable)

	err := r.db.Get(&item, query, userId, itemId)

	return item, err
}

func (r *TodoItemPostgres) DeleteItemById(userId int, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li %s ul  
			WHERE ti.id=$1 AND li.item_id = ti.id AND li.list_id = ul.list_id AND ul.user_id = $2`, todoItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, itemId, userId)

	return err
}

func (r *TodoItemPostgres) UpdateItem(userId int, itemId int, input todo.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf(`title=$%d`, argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf(`description=$%d`, argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf(`done=$%d`, argId))
		args = append(args, *input.Done)
		argId++
	}
	if input.Importance != nil {
		setValues = append(setValues, fmt.Sprintf(`importance=$%d`, argId))
		args = append(args, *input.Importance)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s ul, %s li 
							WHERE li.item_id = ti.id AND ul.list_id = li.list_id AND ti.id=$%d AND ul.user_id=$%d`,
		todoItemsTable, setQuery, usersListsTable, listsItemsTable, argId, argId+1)
	args = append(args, itemId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
