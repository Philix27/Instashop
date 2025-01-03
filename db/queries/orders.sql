-- name: Orders_GetAll :many
SELECT * FROM orders
ORDER BY created_at DESC;

-- name: Orders_GetAllByUserId :many
SELECT * FROM orders
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: Orders_GetById :one
SELECT * FROM orders
WHERE id = $1;

-- name: Orders_UpdateStatus :exec
UPDATE orders
  SET order_status = $2
WHERE id = $1;

-- name: Orders_Create :one
INSERT INTO orders (
  user_id,
  order_status
) VALUES (
  $1,
  "CREATED"
)
RETURNING *;
