package sqlite3

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"asdf/internal/apps/users/internal/application/ports"
	"asdf/internal/dtos"
	"asdf/internal/lib"
)

type Sqlite3Repository struct {
	db     *sql.DB
	logger lib.ILogger
}

var _ ports.IUsersRepository = (*Sqlite3Repository)(nil)

func NewRepository(db *sql.DB, logger lib.ILogger) *Sqlite3Repository {
	return &Sqlite3Repository{
		db,
		logger,
	}
}

const getUserStmt = `
SELECT
	id,
	username,
	email,
	created_at,
	updated_at
FROM
	users
WHERE
	id = ?
LIMIT
	1
`

func (t Sqlite3Repository) GetUser(ctx context.Context, arg *dtos.GetUserParams) (user *dtos.User, err error) {
	if arg.ID <= 0 {
		return nil, lib.ErrNoRecord
	}
	row := t.db.QueryRowContext(ctx, getUserStmt, arg.ID)
	user = &dtos.User{}
	err = row.Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, lib.ErrNoRecord

		default:
			return nil, err
		}
	}
	return user, nil
}

const listUsersStmt = `
SELECT
	id,
	username,
	email,
	created_at,
	updated_at
FROM
	users
WHERE
	id = ?
`

func (t Sqlite3Repository) FindUsers(ctx context.Context, arg *dtos.FindUsersParams) (users []*dtos.User, err error) {
	rows, err := t.db.QueryContext(ctx, listUsersStmt)
	if err != nil {
		t.logger.Error("unable to query rows", "FindUsers", err)
		return nil, err
	}
	defer func() {
		rowsErr := rows.Close()
		t.logger.Error("unable to close rows", "FindUsers", rowsErr)
		err = errors.Join(err, rowsErr)
	}()

	users = []*dtos.User{}
	for rows.Next() {
		user := &dtos.User{}
		if rowErr := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); rowErr != nil {
			t.logger.Error("rows Scan() got errors", "Findusers", rowErr)
			err = errors.Join(err, rowErr)
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

const createUserStmt = `INSERT INTO
	users (username, email, hashed_password)
VALUES
	(?, ?, ?) RETURNING id`

func (t Sqlite3Repository) CreateUser(ctx context.Context, arg *dtos.CreateUserParams) (id int64, err error) {
	row := t.db.QueryRowContext(ctx, createUserStmt, arg.Username, arg.Email, arg.Password)
	err = row.Scan(&id)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "UNIQUE"):
			return -1, lib.ErrNotUnique{
				Msg: err.Error(),
			}

		default:
			return -1, err
		}
	}
	return id, err
}

const updateUserStmt = `
UPDATE
	users
set
	username = ?,
	email = ?,
	updated_at = ?
WHERE
	id = ?
`

func (t Sqlite3Repository) UpdateUser(ctx context.Context, arg *dtos.UpdateUserParams) error {
	if arg.ID <= 0 {
		return lib.ErrNoRecord
	}
	_, err := t.db.ExecContext(ctx, updateUserStmt, arg.Username, arg.Email, lib.ITime(time.Now()).Time(), arg.ID)
	return err
}

const deleteUserStmt = `
DELETE FROM
	users
WHERE
	id = ?
`

func (t Sqlite3Repository) DeleteUser(ctx context.Context, arg *dtos.DeleteUserParams) error {
	if arg.ID <= 0 {
		return lib.ErrNoRecord
	}
	_, err := t.db.ExecContext(ctx, deleteUserStmt, arg.ID)
	return err
}
