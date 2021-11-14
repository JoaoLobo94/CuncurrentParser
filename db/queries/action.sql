-- name: GetAction :one
SELECT * FROM actions
WHERE id = $1 LIMIT 1;

-- name: ListActions :many
SELECT * FROM actions
ORDER BY id;

-- name: CreateAction :one
INSERT INTO actions (
  amount, user_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteActions :exec
DELETE FROM actions;