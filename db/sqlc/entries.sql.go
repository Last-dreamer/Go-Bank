// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: entries.sql

package db

import (
	"context"
	"database/sql"
)

const createEntries = `-- name: CreateEntries :one
INSERT INTO entries (
    account_id,
    money
) VALUES (
    $1, $2
) RETURNING id, account_id, money, created_at
`

type CreateEntriesParams struct {
	AccountID sql.NullInt64 `json:"account_id"`
	Money     int64         `json:"money"`
}

func (q *Queries) CreateEntries(ctx context.Context, arg CreateEntriesParams) (Entries, error) {
	row := q.queryRow(ctx, q.createEntriesStmt, createEntries, arg.AccountID, arg.Money)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Money,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM entries
Where id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteEntryStmt, deleteEntry, id)
	return err
}

const getEntries = `-- name: GetEntries :many
SELECT id, account_id, money, created_at FROM entries
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetEntriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetEntries(ctx context.Context, arg GetEntriesParams) ([]Entries, error) {
	rows, err := q.query(ctx, q.getEntriesStmt, getEntries, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entries
	for rows.Next() {
		var i Entries
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Money,
			&i.CreatedAt,
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

const getEntry = `-- name: GetEntry :one
SELECT id, account_id, money, created_at FROM entries
Where id = $1 LIMIT $1
`

func (q *Queries) GetEntry(ctx context.Context, limit int32) (Entries, error) {
	row := q.queryRow(ctx, q.getEntryStmt, getEntry, limit)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Money,
		&i.CreatedAt,
	)
	return i, err
}

const updateEntry = `-- name: UpdateEntry :one
UPDATE entries
SET money = $2
Where id = $1
RETURNING id, account_id, money, created_at
`

type UpdateEntryParams struct {
	ID    int64 `json:"id"`
	Money int64 `json:"money"`
}

func (q *Queries) UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entries, error) {
	row := q.queryRow(ctx, q.updateEntryStmt, updateEntry, arg.ID, arg.Money)
	var i Entries
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Money,
		&i.CreatedAt,
	)
	return i, err
}