-- name: CreateCbtlog :one
INSERT INTO
    cbtlogs (filename, relative_path, url)
VALUES (?, ?, ?) RETURNING *;

-- name: DeleteCbtlogByFilename :one
DELETE FROM cbtlogs
WHERE filename = ?
RETURNING *;

-- name: UpdateCbtlogUploadStatus :exec
UPDATE cbtlogs
SET
    upload_status = ?,
    upload_status_reason = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?;

-- name: UpdateCbtlogUploadResult :exec
UPDATE cbtlogs
SET
    url = ?,
encounter_success = ?,
challenge_mode = ?,
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
    AND (
        substr(
            filename,
            1,
            length(sqlc.narg ('date'))
        ) = sqlc.narg ('date')
        OR sqlc.narg ('date') IS NULL
    )
    AND (
        (
            substr(filename, 1, 8) >= sqlc.narg ('from_date')
            OR sqlc.narg ('from_date') IS NULL
        )
    )
    AND (
        (
            substr(filename, 1, 8) <= sqlc.narg ('to_date')
            OR sqlc.narg ('to_date') IS NULL
        )
    )
    AND (
        challenge_mode = sqlc.narg ('challenge_mode')
	OR sqlc.narg ('challenge_mode') IS NULL
    )
    AND (
        encounter_success = sqlc.narg ('encounter_success ')
	OR sqlc.narg ('encounter_success ') IS NULL
    )
ORDER BY created_at DESC;

-- name: GetCbtlogByFileName :one
SELECT * FROM cbtlogs WHERE filename = ?;
