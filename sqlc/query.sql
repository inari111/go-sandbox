-- name: Get :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: Create :one
INSERT INTO users (
  id, name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: Update :exec
UPDATE users SET name = $2
WHERE id = $1;

-- name: Delete :exec
DELETE FROM users WHERE id = $1;

-- name: List :many
SELECT * FROM users
ORDER BY id;

-- name: Count :one
SELECT count(*) FROM users;
