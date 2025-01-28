package sqlite3

import (
	"context"
	"database/sql"
	"errors"

	"with-alias/internal/apps/scheduler/internal/application/ports"
	"with-alias/internal/dtos"
	"with-alias/internal/lib"
)

type Sqlite3Repository struct {
	db     *sql.DB
	logger lib.ILogger
}

var _ ports.ISchedulerRepository = (*Sqlite3Repository)(nil)

func NewRepository(db *sql.DB, logger lib.ILogger) *Sqlite3Repository {
	return &Sqlite3Repository{
		db,
		logger,
	}
}

const getSchedulerStmt = `
SELECT
	id
FROM
	scheduler
WHERE
	id = ?
LIMIT
	1
`

func (t Sqlite3Repository) GetScheduler(ctx context.Context, arg *dtos.GetSchedulerParams) (scheduler *dtos.Scheduler, err error) {
	if arg.ID <= 0 {
		return nil, lib.ErrNoRecord
	}
	row := t.db.QueryRowContext(ctx, getSchedulerStmt, arg.ID)
	scheduler = &dtos.Scheduler{}
	err = row.Scan(
		&scheduler.ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, lib.ErrNoRecord
		}
	}
	return scheduler, nil
}

const listSchedulerStmt = `
SELECT
	id
FROM
	scheduler
`

func (t Sqlite3Repository) FindScheduler(ctx context.Context, arg *dtos.FindSchedulerParams) (scheduler []*dtos.Scheduler, err error) {
	rows, err := t.db.QueryContext(ctx, listSchedulerStmt)
	if err != nil {
		return nil, err
	}
	defer func() {
		rowsErr := rows.Close()
		if err == nil {
			err = rowsErr
		} else {
			t.logger.Error("unable to close rows", "FindScheduler", err)
		}
	}()
	items := []*dtos.Scheduler{}
	for rows.Next() {
		i := &dtos.Scheduler{}
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

const createSchedulerStmt = `INSERT INTO
	scheduler (task, description)
VALUES
	(?, ?) RETURNING id`

func (t Sqlite3Repository) CreateScheduler(ctx context.Context, arg *dtos.CreateSchedulerParams) (id int64, err error) {
	row := t.db.QueryRowContext(ctx, createSchedulerStmt, arg.Task, arg.Description)
	err = row.Scan(&id)
	return id, err
}

const updateSchedulerStmt = `
UPDATE
	scheduler
set
	task = ?,
	description = ?
WHERE
	id = ?
`

func (t Sqlite3Repository) UpdateScheduler(ctx context.Context, arg *dtos.UpdateSchedulerParams) error {
	if arg.ID <= 0 {
		return lib.ErrNoRecord
	}
	_, err := t.db.ExecContext(ctx, updateSchedulerStmt, arg.Task, arg.Description, arg.ID)
	return err
}

const deleteSchedulerStmt = `
DELETE FROM
	scheduler
WHERE
	id = ?
`

func (t Sqlite3Repository) DeleteScheduler(ctx context.Context, arg *dtos.DeleteSchedulerParams) error {
	if arg.ID <= 0 {
		return lib.ErrNoRecord
	}
	_, err := t.db.ExecContext(ctx, deleteSchedulerStmt, arg.ID)
	return err
}
