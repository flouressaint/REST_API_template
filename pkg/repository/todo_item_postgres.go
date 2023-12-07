package repository

import (
	"REST_API"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item REST_API.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf(
		`INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id`,
		todoItemsTable,
	)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err := row.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createListsItemQuery := fmt.Sprintf(
		`INSERT INTO %s (list_id, item_id) VALUES ($1, $2)`,
		listsItemsTable,
	)
	_, err = tx.Exec(createListsItemQuery, listId, itemId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(listId int) ([]REST_API.TodoItem, error) {
	var items []REST_API.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
							INNER JOIN %s li ON li.id = li.list_id 
							WHERE li.list_id = $1`,
		todoItemsTable,
		listsItemsTable,
	)
	err := r.db.Select(&items, query, listId)

	return items, err
}
