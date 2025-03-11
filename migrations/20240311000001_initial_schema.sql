-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    description TEXT,
    join_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ponds (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    member_count INTEGER DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ripples (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    like_count INTEGER DEFAULT 0,
    comment_count INTEGER DEFAULT 0,
    author_username TEXT NOT NULL DEFAULT '',
    pond_name TEXT NOT NULL DEFAULT 'OFFICIAL',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_ponds (
    user_id INTEGER NOT NULL REFERENCES users(id),
    pond_id INTEGER NOT NULL REFERENCES ponds(id),
    joined_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, pond_id)
);

CREATE INDEX idx_ripples_author ON ripples(author_username);
CREATE INDEX idx_ripples_pond ON ripples(pond_name);
CREATE INDEX idx_user_ponds_user ON user_ponds(user_id);
CREATE INDEX idx_user_ponds_pond ON user_ponds(pond_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_ripples_author;
DROP INDEX IF EXISTS idx_ripples_pond;
DROP INDEX IF EXISTS idx_user_ponds_user;
DROP INDEX IF EXISTS idx_user_ponds_pond;

DROP TABLE IF EXISTS user_ponds;
DROP TABLE IF EXISTS ripples;
DROP TABLE IF EXISTS ponds;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd 