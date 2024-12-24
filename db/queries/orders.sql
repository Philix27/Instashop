-- name: Orders_GetAll :many
SELECT * FROM orders
ORDER BY created_at DESC;

-- name: Orders_GetAllByUserId :many
SELECT * FROM orders
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: Orders_Update :exec
UPDATE orders
  SET order_status = $2
WHERE id = $1;