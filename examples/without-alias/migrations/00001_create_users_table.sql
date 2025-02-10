-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id integer primary key,
	username text NOT NULL,
	hashed_password text NOT NULL,
	role text default 'level_1',
	email text NOT NULL,
	created_at datetime default CURRENT_TIMESTAMP,
	updated_at datetime default CURRENT_TIMESTAMP,
	UNIQUE(username),
	UNIQUE(email),
	CHECK(
		length(username) >= 2
		AND length(email) >= 4
		AND length(hashed_password) >= 1
	)
);

CREATE UNIQUE INDEX idx_users_email ON users(email);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

-- +goose StatementEnd