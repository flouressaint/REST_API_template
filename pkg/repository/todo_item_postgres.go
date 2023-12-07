package repository

import (
	"REST_API"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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
	if err := r.db.Select(&items, query, listId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (REST_API.TodoItem, error) {
	var item REST_API.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti
							INNER JOIN %s li ON li.id = li.item_id
							INNER JOIN %s ul ON ul.list_id = li.list_id
							WHERE ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)
	if err := r.db.Get(&item, query, userId, itemId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TodoItemPostgres) Update(userId, itemId int, input REST_API.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`Update %s ti SET %s
							FROM %s li
							INNER JOIN %s ul
							ON li.list_id = ul.list_id
							WHERE ti.id = li.item_id
								AND li.item_id=$%d
								AND ul.user_id=$%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, itemId, userId)

	logrus.Debug("UpdateQuery: %s", query)
	logrus.Debug("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul
							WHERE ti.id = li.item_id AND
								  li.list_id = ul.list_id AND
								  ul.user_id = $1 AND
								  ti.id = $2`,
		todoItemsTable,
		listsItemsTable,
		usersListsTable,
	)
	_, err := r.db.Exec(query, userId, itemId)

	return err
}
