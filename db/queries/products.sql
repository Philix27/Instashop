-- name: Products_GetAll :many
SELECT * FROM products
ORDER BY created_at DESC;

-- name: Products_GetOne :one
SELECT * FROM products
WHERE id = $1;


-- name: Products_Create :one
INSERT INTO products (
  title,
  description,
  price,
  stock,
  image_url
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING *;

-- name: Product_Update :exec
UPDATE products
  SET title = $2,
  description = $3,
  stock = $4,
  image_url = $5,
  price = $6
WHERE id = $1;

-- name: Product_Delete :exec
DELETE FROM products
WHERE id = $1;