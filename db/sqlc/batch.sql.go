// Code generated by sqlc. DO NOT EDIT.
// source: batch.sql

package db

import (
	"context"
)

const createBatch = `-- name: CreateBatch :one
INSERT INTO batches (
  dispatched, amount, user_id
) VALUES (
  $1, $2, $3
)
RETURNING id, dispatched, amount, user_id, created_at
`

type CreateBatchParams struct {
	Dispatched bool
	Amount     float64
	UserID     int32
}

func (q *Queries) CreateBatch(ctx context.Context, arg CreateBatchParams) (Batch, error) {
	row := q.db.QueryRowContext(ctx, createBatch, arg.Dispatched, arg.Amount, arg.UserID)
	var i Batch
	err := row.Scan(
		&i.ID,
		&i.Dispatched,
		&i.Amount,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteBatches = `-- name: DeleteBatches :exec
DELETE FROM batches
`

func (q *Queries) DeleteBatches(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteBatches)
	return err
}

const getBatch = `-- name: GetBatch :one
SELECT id, dispatched, amount, user_id, created_at FROM batches
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBatch(ctx context.Context, id int32) (Batch, error) {
	row := q.db.QueryRowContext(ctx, getBatch, id)
	var i Batch
	err := row.Scan(
		&i.ID,
		&i.Dispatched,
		&i.Amount,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const listBatches = `-- name: ListBatches :many
SELECT id, dispatched, amount, user_id, created_at FROM batches
ORDER BY id
`

func (q *Queries) ListBatches(ctx context.Context) ([]Batch, error) {
	rows, err := q.db.QueryContext(ctx, listBatches)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Batch
	for rows.Next() {
		var i Batch
		if err := rows.Scan(
			&i.ID,
			&i.Dispatched,
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

const listDispatchedBatches = `-- name: ListDispatchedBatches :many
SELECT id, dispatched, amount, user_id, created_at FROM batches
WHERE dispatched = $1
ORDER BY id
`

func (q *Queries) ListDispatchedBatches(ctx context.Context, dispatched bool) ([]Batch, error) {
	rows, err := q.db.QueryContext(ctx, listDispatchedBatches, dispatched)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Batch
	for rows.Next() {
		var i Batch
		if err := rows.Scan(
			&i.ID,
			&i.Dispatched,
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

const updateBatch = `-- name: UpdateBatch :exec
UPDATE batches 
SET amount= $2, dispatched= $3
WHERE id = $1
`

type UpdateBatchParams struct {
	ID         int32
	Amount     float64
	Dispatched bool
}

func (q *Queries) UpdateBatch(ctx context.Context, arg UpdateBatchParams) error {
	_, err := q.db.ExecContext(ctx, updateBatch, arg.ID, arg.Amount, arg.Dispatched)
	return err
}
