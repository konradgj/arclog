-- +goose Up
CREATE TABLE uploads (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    file_path TEXT NOT NULL UNIQUE,
    url TEXT UNIQUE,
    status TEXT NOT NULL DEFAULT 'pending',
    status_reason TEXT DEFAULT 'create',
    active INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE uploads;