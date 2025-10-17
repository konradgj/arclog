-- name: CreateCbtlog :one
INSERT INTO
    cbtlogs (filename, relative_path, url)
VALUES (?, ?, ?) RETURNING *;

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

-- name: ListCbtlogsByFilters :many
SELECT *
FROM cbtlogs
WHERE (
        upload_status = sqlc.narg ('upload_status')
        OR sqlc.narg ('upload_status') IS NULL
    )
    AND (
        relative_path LIKE COALESCE(
            sqlc.narg ('relative_path'),
            relative_path
        )
    )
ORDER BY created_at DESC;

-- name: GetCbtlogByFileName :one
SELECT * FROM cbtlogs WHERE filename = ?;