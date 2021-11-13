-- name: GetBank_Transaction :one
SELECT * FROM bank_transactions
WHERE id = $1 LIMIT 1;

-- name: ListBank_Transactions :many
SELECT * FROM bank_transactions
ORDER BY name;

-- name: CreateBank_Transaction :one
INSERT INTO bank_transactions (
  amount, user_id
) VALUES (
  $1, $2
)
RETURNING *;