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

-- name: User_Update :exec
UPDATE users
  SET first_name = $2,
  last_name = $3,
  middle_name = $4,
  dob = $5
WHERE id = $1;

-- name: User_PhoneNumber :exec
UPDATE users
  SET phone = $2
WHERE id = $1;

-- name: User_ResetPassword :exec
UPDATE users
  SET hashed_password = $2
WHERE id = $1;

-- name: Kyc_UpdateCredentials :exec
UPDATE users
  SET first_name = $2,
  last_name = $3,
  phone = $4
WHERE id = $1;