-- name: ListTransactions :many
SELECT * FROM transactions
ORDER BY id;

-- name: CreateTransaction :one
INSERT INTO transactions (
  amount, user_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteTransactions :exec
DELETE FROM transactions;
