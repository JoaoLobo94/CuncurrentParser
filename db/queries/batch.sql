-- name: GetBatch :one
SELECT * FROM batches
WHERE id = $1 LIMIT 1;

-- name: ListBatches :many
SELECT * FROM batches
ORDER BY id;

-- name: CreateBatch :one
INSERT INTO batches (
  dispatched, amount, user_id
) VALUES (
  $1, $2, $3
)
RETURNING *;