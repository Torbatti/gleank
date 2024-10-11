// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package models

import (
	"context"
	"database/sql"
)

const createFolder = `-- name: CreateFolder :one
INSERT INTO folders (
  user, url , public , name , description
) VALUES (
  ?, ? , ? , ? , ?
)
RETURNING id, url, name, description, user, public
`

type CreateFolderParams struct {
	User        int64
	Url         string
	Public      sql.NullBool
	Name        string
	Description sql.NullString
}

func (q *Queries) CreateFolder(ctx context.Context, arg CreateFolderParams) (Folder, error) {
	row := q.db.QueryRowContext(ctx, createFolder,
		arg.User,
		arg.Url,
		arg.Public,
		arg.Name,
		arg.Description,
	)
	var i Folder
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Name,
		&i.Description,
		&i.User,
		&i.Public,
	)
	return i, err
}

const createLink = `-- name: CreateLink :one
INSERT INTO links (
  folder, url , name , description
) VALUES (
  ?, ? , ? , ?
)
RETURNING id, url, name, description, folder
`

type CreateLinkParams struct {
	Folder      int64
	Url         string
	Name        sql.NullString
	Description sql.NullString
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) (Link, error) {
	row := q.db.QueryRowContext(ctx, createLink,
		arg.Folder,
		arg.Url,
		arg.Name,
		arg.Description,
	)
	var i Link
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Name,
		&i.Description,
		&i.Folder,
	)
	return i, err
}

const deleteFolder = `-- name: DeleteFolder :exec
DELETE FROM folders
WHERE id = ?
`

func (q *Queries) DeleteFolder(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFolder, id)
	return err
}

const deleteLink = `-- name: DeleteLink :exec
DELETE FROM links
WHERE id = ?
`

func (q *Queries) DeleteLink(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteLink, id)
	return err
}

const getFolder = `-- name: GetFolder :one
SELECT id, url, name, description, user, public FROM folders
WHERE id = ? LIMIT 1
`

func (q *Queries) GetFolder(ctx context.Context, id int64) (Folder, error) {
	row := q.db.QueryRowContext(ctx, getFolder, id)
	var i Folder
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Name,
		&i.Description,
		&i.User,
		&i.Public,
	)
	return i, err
}

const getLink = `-- name: GetLink :one
SELECT id, url, name, description, folder FROM links
WHERE id = ? LIMIT 1
`

func (q *Queries) GetLink(ctx context.Context, id int64) (Link, error) {
	row := q.db.QueryRowContext(ctx, getLink, id)
	var i Link
	err := row.Scan(
		&i.ID,
		&i.Url,
		&i.Name,
		&i.Description,
		&i.Folder,
	)
	return i, err
}

const listFolders = `-- name: ListFolders :many
SELECT id, url, name, description, user, public FROM folders
WHERE id = ?
ORDER BY name
`

func (q *Queries) ListFolders(ctx context.Context, id int64) ([]Folder, error) {
	rows, err := q.db.QueryContext(ctx, listFolders, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Folder
	for rows.Next() {
		var i Folder
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.Name,
			&i.Description,
			&i.User,
			&i.Public,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLinks = `-- name: ListLinks :many
SELECT id, url, name, description, folder FROM links
WHERE id = ?
ORDER BY name
`

func (q *Queries) ListLinks(ctx context.Context, id int64) ([]Link, error) {
	rows, err := q.db.QueryContext(ctx, listLinks, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.ID,
			&i.Url,
			&i.Name,
			&i.Description,
			&i.Folder,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFolder = `-- name: UpdateFolder :exec
UPDATE folders
set name = ?,
description = ?,
public = ?
WHERE id = ?
`

type UpdateFolderParams struct {
	Name        string
	Description sql.NullString
	Public      sql.NullBool
	ID          int64
}

func (q *Queries) UpdateFolder(ctx context.Context, arg UpdateFolderParams) error {
	_, err := q.db.ExecContext(ctx, updateFolder,
		arg.Name,
		arg.Description,
		arg.Public,
		arg.ID,
	)
	return err
}

const updateLink = `-- name: UpdateLink :exec
UPDATE links
set url = ?,
name = ?,
description = ?
WHERE id = ?
`

type UpdateLinkParams struct {
	Url         string
	Name        sql.NullString
	Description sql.NullString
	ID          int64
}

func (q *Queries) UpdateLink(ctx context.Context, arg UpdateLinkParams) error {
	_, err := q.db.ExecContext(ctx, updateLink,
		arg.Url,
		arg.Name,
		arg.Description,
		arg.ID,
	)
	return err
}
