package postgres

import (
	"context"
	"database/sql"
	"errors"

	"without-alias/internal/apps/users/internal/application/ports"
	"without-alias/internal/dtos"
	"without-alias/internal/lib"
)

type PostgresRepository struct {
	db     *sql.DB
	logger lib.ILogger
}

var _ ports.IUsersRepository = (*PostgresRepository)(nil)

func NewRepository(db *sql.DB, logger lib.ILogger) *PostgresRepository {
	return &PostgresRepository{
		db,
		logger,
	}
}

const getUserStmt = `
SELECT
	id
FROM
	users
WHERE
	id = ?
LIMIT
	1
`

func (t PostgresRepository) GetUser(ctx context.Context, arg *dtos.GetUserParams) (user *dtos.User, err error) {
	if arg.ID <= 0 {
		return nil, lib.ErrNoRecord
	}
	row := t.db.QueryRowContext(ctx, getUserStmt, arg.ID)
	user = &dtos.User{}
	err = row.Scan(
		&user.ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, lib.ErrNoRecord
		}
	}
	return user, nil
}

const listUsersStmt = `
SELECT
	id
FROM
	users
`

func (t PostgresRepository) FindUsers(ctx context.Context, arg *dtos.FindUsersParams) (users []*dtos.User, err error) {
	rows, err := t.db.QueryContext(ctx, listUsersStmt)
	if err != nil {
		return nil, err
	}
	defer func() {
		rowsErr := rows.Close()
		if err == nil {
			err = rowsErr
		} else {
			t.logger.Error("unable to close rows", "FindUsers", err)
		}
	}()
	items := []*dtos.User{}
	for rows.Next() {
		i := &dtos.User{}
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

const createUserStmt = `INSERT INTO
	users (task, description)
VALUES
	(?, ?) RETURNING id`

func (t PostgresRepository) CreateUser(ctx context.Context, arg *dtos.CreateUserParams) (id int64, err error) {
	row := t.db.QueryRowContext(ctx, createUserStmt, arg.Task, arg.Description)
	err = row.Scan(&id)
	return id, err
}

const updateUserStmt = `
UPDATE
	users
set
	task = ?,
	description = ?
WHERE
	id = ?
`

func (t PostgresRepository) UpdateUser(ctx context.Context, arg *dtos.UpdateUserParams) error {
	if arg.ID <= 0 {
		return lib.ErrNoRecord
	}
	_, err := t.db.ExecContext(ctx, updateUserStmt, arg.Task, arg.Description, arg.ID)
	return err
}

const deleteUserStmt = `
DELETE FROM
	users
WHERE
	id = ?
`

func (t PostgresRepository) DeleteUser(ctx context.Context, arg *dtos.DeleteUserParams) error {
	if arg.ID <= 0 {
		return lib.ErrNoRecord
	}
	_, err := t.db.ExecContext(ctx, deleteUserStmt, arg.ID)
	return err
}
