-- name: OrderItem_GetAll :many
SELECT * FROM order_items
ORDER BY created_at DESC;

-- name: OrderItem_GetAllOrderId :many
SELECT * FROM order_items
WHERE order_id = $1
ORDER BY created_at DESC;

-- name: OrderItem_IncrementQuantity :exec
UPDATE order_items
  SET quantity = $2
WHERE id = $1;

-- name: OrderItem_DecrementQuantity :exec
UPDATE order_items
  SET quantity = $2
WHERE id = $1;


-- name: OrderItem_DeleteItem :exec
DELETE FROM order_items
WHERE order_id = $1 AND product_id = $2;


-- name: OrderItem_ClearOrders :exec
DELETE FROM order_items
WHERE order_id = 1;
