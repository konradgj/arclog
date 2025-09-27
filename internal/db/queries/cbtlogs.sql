-- name: CreateCbtlog :one
INSERT INTO cbtlogs (file_path, url) VALUES (?, ?) RETURNING *;

-- name: UpdateCtblogUploadStatus :exec
UPDATE cbtlogs
SET
    upload_status = ?,
    upload_status_reason = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?;

-- name: UpdateCbtlogUrl :exec
UPDATE cbtlogs
SET
    url = ?,
    upload_status = ?,
    upload_status_reason = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?;

-- name: ListCbtlogsByUploadStatus :many
SELECT *
FROM cbtlogs
WHERE
    upload_status = ?
ORDER BY created_at DESC;