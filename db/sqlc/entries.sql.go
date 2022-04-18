// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: entries.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (account_id, amount)
VALUES (?, ?);
`
type CreateEntryArgs struct {
	Account_id int64 `json:"account_id" binding:"required"`
	Amount int64`json:"amount" binding:"required"`
}

func (q *Queries) CreateEntry(ctx context.Context, args CreateEntryArgs) (Entry, error) {
	r, err := q.db.ExecContext(ctx, createEntry, args.Account_id, args.Amount)
	var i Entry
	if err != nil {
		return i, err
	}
	lastInsertedId, err := r.LastInsertId()
	if err != nil {
		return i, err
	}
	row := q.db.QueryRowContext(ctx, getLastInserted("entries", lastInsertedId))
	err = row.Scan(
		&i.ID,
		&i.AccountID,
		&i.CreatedAt,
		&i.Amount,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = ?
`

func (q *Queries) DeleteEntry(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntry, id)
	return err
}

const getEntries = `-- name: GetEntries :many
SELECT id, account_id, created_at, amount FROM entries LIMIT ?, ?`

type GetEntriesArgs struct {
	Limit int
	Offset int
}

func (q *Queries) GetEntries(ctx context.Context, args GetEntriesArgs) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, getEntries, args.Limit, args.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.CreatedAt,
			&i.Amount,
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
SELECT id, account_id, created_at, amount FROM entries
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.CreatedAt,
		&i.Amount,
	)
	return i, err
}

const updateEntryAmount = `-- name: UpdateEntryAmount :one
UPDATE entries SET amount = ? WHERE id = ?
`

type UpdateEntryArgs struct {
	Amount int64 `json:"amount" binding:"required"`
	Id int64 `json:"id" binding:"required"`
} 

func (q *Queries) UpdateEntryAmount(ctx context.Context, args UpdateEntryArgs) (Entry, error) {
	_, err := q.db.ExecContext(ctx, updateEntryAmount, args.Amount, args.Id)
	var i Entry
	if err != nil {
		return i, err
	}
	row := q.db.QueryRowContext(ctx, getLastInserted("entries", args.Id))
	err = row.Scan(
		&i.ID,
		&i.AccountID,
		&i.CreatedAt,
		&i.Amount,
	)
	return i, err
}
