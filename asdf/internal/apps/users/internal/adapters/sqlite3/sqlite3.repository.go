package sqlite3

import (
	"context"
	"database/sql"
	"errors"

	"asdf/internal/apps/users/internal/application/ports"
	"asdf/internal/dtos"
	"asdf/internal/lib"
)

type Sqlite3Repository struct {
	db *sql.DB
}

var _ ports.IUsersRepository = (*Sqlite3Repository)(nil)

func NewRepository(db *sql.DB) *Sqlite3Repository {
	return &Sqlite3Repository{
		db,
	}
}

func (t Sqlite3Repository) GetUsers(ctx context.Context, arg *dtos.GetUsersParams) (user *dtos.User, err error) {
	if arg.ID < 1 {
		return nil, lib.ErrNoRecord
	}
	const stmt = `SELECT ID from users 
    				WHERE id = ?`
	row := t.db.QueryRow(stmt, arg.ID)

	u := &dtos.User{}

	err = row.Scan(&u.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, lib.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return u, nil
}

func (t Sqlite3Repository) FindUsers(ctx context.Context, arg *dtos.FindUsersParams) ([]*dtos.User, error) {
	return nil, nil
}

func (t Sqlite3Repository) CreateUsers(ctx context.Context, arg *dtos.CreateUsersParams) (lastInsertId int64, err error) {
	const stmt string = `INSERT INTO users (col1, col2, col3, col3)
							VALUES(?, ?, ?, ?)`
	result, err := t.db.Exec(stmt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (t Sqlite3Repository) UpdateUsers(ctx context.Context, arg *dtos.UpdateUsersParams) error {
	return nil
}

func (t Sqlite3Repository) DeleteUsers(ctx context.Context, arg *dtos.DeleteUsersParams) error {
	return nil
}
