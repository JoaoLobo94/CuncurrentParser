// Code generated by sqlc. DO NOT EDIT.
// source: action.sql

package db

import (
	"context"
)

const createAction = `-- name: CreateAction :one
INSERT INTO actions (
  amount, user_id
) VALUES (
  $1, $2
)
RETURNING id, amount, user_id, created_at
`

type CreateActionParams struct {
	Amount float64
	UserID int32
}

func (q *Queries) CreateAction(ctx context.Context, arg CreateActionParams) (Action, error) {
	row := q.db.QueryRowContext(ctx, createAction, arg.Amount, arg.UserID)
	var i Action
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const getAction = `-- name: GetAction :one
SELECT id, amount, user_id, created_at FROM actions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAction(ctx context.Context, id int32) (Action, error) {
	row := q.db.QueryRowContext(ctx, getAction, id)
	var i Action
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const listActions = `-- name: ListActions :many
SELECT id, amount, user_id, created_at FROM actions
ORDER BY id
`

func (q *Queries) ListActions(ctx context.Context) ([]Action, error) {
	rows, err := q.db.QueryContext(ctx, listActions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Action
	for rows.Next() {
		var i Action
		if err := rows.Scan(
			&i.ID,
			&i.Amount,
			&i.UserID,
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