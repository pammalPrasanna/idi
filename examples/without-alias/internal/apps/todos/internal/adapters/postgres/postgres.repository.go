package postgres

import (
	"context"
	"database/sql"
	"errors"

	"without-alias/internal/apps/todos/internal/application/ports"
	"without-alias/internal/dtos"
	"without-alias/internal/lib"
)

type PostgresRepository struct {
	db     *sql.DB
	logger lib.ILogger
}

var _ ports.ITodosRepository = (*PostgresRepository)(nil)

func NewRepository(db *sql.DB, logger lib.ILogger) *PostgresRepository {
	return &PostgresRepository{
		db,
		logger,
	}
}

const getTodoStmt = `
SELECT
	id
FROM
	todos
WHERE
	id = ?
LIMIT
	1
`

func (t PostgresRepository) GetTodo(ctx context.Context, arg *dtos.GetTodoParams) (todo *dtos.Todo, err error) {
	if arg.ID <= 0 {
		return nil, lib.ErrNoRecord
	}
	row := t.db.QueryRowContext(ctx, getTodoStmt, arg.ID)
	todo = &dtos.Todo{}
	err = row.Scan(
		&todo.ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, lib.ErrNoRecord
		}
	}
	return todo, nil
}

const listTodosStmt = `
SELECT
	id
FROM
	todos
`

func (t PostgresRepository) FindTodos(ctx context.Context, arg *dtos.FindTodosParams) (todos []*dtos.Todo, err error) {
	rows, err := t.db.QueryContext(ctx, listTodosStmt)
	if err != nil {
		return nil, err
	}
	defer func() {
		rowsErr := rows.Close()
		if err == nil {
			err = rowsErr
		} else {
			t.logger.Error("unable to close rows", "FindTodos", err)
		}
	}()
	items := []*dtos.Todo{}
	for rows.Next() {
		i := &dtos.Todo{}
		if err := rows.Scan(
			&i.ID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createTodoStmt = `INSERT INTO
	todos (task, description)
VALUES
	(?, ?) RETURNING id`

func (t PostgresRepository) CreateTodo(ctx context.Context, arg *dtos.CreateTodoParams) (id int64, err error) {
	row := t.db.QueryRowContext(ctx, createTodoStmt, arg.Task, arg.Description)
	err = row.Scan(&id)
	return id, err
}

const updateTodoStmt = `
UPDATE
	todos
set
	task = ?,
	description = ?
WHERE
	id = ?
`

func (t PostgresRepository) UpdateTodo(ctx context.Context, arg *dtos.UpdateTodoParams) error {
	if arg.ID <= 0 {
		return lib.ErrNoRecord
	}
	_, err := t.db.ExecContext(ctx, updateTodoStmt, arg.Task, arg.Description, arg.ID)
	return err
}

const deleteTodoStmt = `
DELETE FROM
	todos
WHERE
	id = ?
`

func (t PostgresRepository) DeleteTodo(ctx context.Context, arg *dtos.DeleteTodoParams) error {
	if arg.ID <= 0 {
		return lib.ErrNoRecord
	}
	_, err := t.db.ExecContext(ctx, deleteTodoStmt, arg.ID)
	return err
}
