-- name: CreateUpload :one
INSERT INTO uploads (file_path, url) VALUES (?, ?) RETURNING *;