-- name: CreateUpload :one
INSERT INTO uploads (file_path, url) VALUES (?, ?) RETURNING *;

-- name: UpdateUploadStatus :exec
UPDATE uploads
SET
    status = ?,
    status_reason = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?;

-- name: UpdateUploadUrl :exec
UPDATE uploads
SET
    url = ?,
    status = ?,
    status_reason = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?;

-- name: ListUploadsByStatus :many
SELECT * FROM uploads WHERE status = ? ORDER BY created_at DESC;