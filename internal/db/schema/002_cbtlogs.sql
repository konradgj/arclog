-- +goose Up
ALTER TABLE cbtlogs ADD COLUMN encounter_success INTEGER;
ALTER TABLE cbtlogs ADD COLUMN challenge_mode INTEGER;

-- +goose Down
CREATE TABLE cbtlogs_new (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    filename TEXT NOT NULL UNIQUE,
    relative_path TEXT,
    url TEXT UNIQUE,
    upload_status TEXT NOT NULL DEFAULT 'pending',
    upload_status_reason TEXT NOT NULL DEFAULT 'create',
    active INTEGER NOT NULL DEFAULT 1,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO cbtlogs_new (
    id,
    filename,
    relative_path,
    url,
    upload_status,
    upload_status_reason,
    active,
    created_at,
    updated_at
)
SELECT
    id,
    filename,
    relative_path,
    url,
    upload_status,
    upload_status_reason,
    active,
    created_at,
    updated_at
FROM cbtlogs;
DROP TABLE cbtlogs;
ALTER TABLE cbtlogs_new RENAME TO cbtlogs;
