// Code generated by sqlc. DO NOT EDIT.
// source: transaction.sql

package db

import (
	"context"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
  amount, user_id
) VALUES (
  $1, $2
)
RETURNING id, amount, user_id, created_at
`

type CreateTransactionParams struct {
	Amount float64
	UserID int32
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction, arg.Amount, arg.UserID)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const listTransactions = `-- name: ListTransactions :many
SELECT id, amount, user_id, created_at FROM transactions
ORDER BY id
`

func (q *Queries) ListTransactions(ctx context.Context) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, listTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
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
