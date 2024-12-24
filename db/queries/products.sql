-- name: Products_GetAll :many
SELECT * FROM products
ORDER BY created_at DESC;


-- name: Products_Create :one
INSERT INTO products (
  title,
  description,
  image_url
) VALUES (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: Product_Update :exec
UPDATE products
  SET title = $2,
  description = $3,
  image_url = $4
WHERE id = $1;