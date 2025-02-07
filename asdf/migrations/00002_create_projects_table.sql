-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects (
	id integer primary key,
	name text,
	description text,
	user_id integer,
	archived boolean,
	created_at datetime default CURRENT_TIMESTAMP,
	updated_at datetime default CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS projects;
-- +goose StatementEnd
