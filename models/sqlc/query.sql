-- name: GetLink :one
SELECT * FROM links
WHERE id = ? LIMIT 1;
-- name: ListLinks :many
SELECT * FROM links
WHERE id = ?
ORDER BY name;
-- name: CreateLink :one
INSERT INTO links (
  folder, url , name , description
) VALUES (
  ?, ? , ? , ?
)
RETURNING *;
-- name: UpdateLink :exec
UPDATE links
set url = ?,
name = ?,
description = ?
WHERE id = ?;
-- name: DeleteLink :exec
DELETE FROM links
WHERE id = ?;


-- name: GetFolder :one
SELECT * FROM folders
WHERE id = ? LIMIT 1;
-- name: ListFolders :many
SELECT * FROM folders
WHERE id = ?
ORDER BY name;
-- name: CreateFolder :one
INSERT INTO folders (
  user, url , public , name , description
) VALUES (
  ?, ? , ? , ? , ?
)
RETURNING *;
-- name: UpdateFolder :exec
UPDATE folders
set name = ?,
description = ?,
public = ?
WHERE id = ?;
-- name: DeleteFolder :exec
DELETE FROM folders
WHERE id = ?;

