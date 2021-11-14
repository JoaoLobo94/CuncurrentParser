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

-- name: UpdateBatch :exec
UPDATE batches 
SET amount= $2, dispatched= $3
WHERE id = $1;

-- name: DeleteBatches :exec
DELETE FROM batches;