-- ! #SECTION user

-- name: User_GetById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: User_GetByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: User_Create :one
INSERT INTO users (
  email,
  hashed_password 
) VALUES (
  $1,
  $2
)
RETURNING *;

-- name: User_ResetPassword :exec
UPDATE users
  SET hashed_password = $2
WHERE id = $1;

