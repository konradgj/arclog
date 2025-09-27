-- +goose Up
CREATE TABLE cbtlogs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    file_path TEXT NOT NULL UNIQUE,
    url TEXT UNIQUE,
    upload_status TEXT NOT NULL DEFAULT 'pending',
    upload_status_reason TEXT NOT NULL DEFAULT 'create',
    active INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE cbtlogs;