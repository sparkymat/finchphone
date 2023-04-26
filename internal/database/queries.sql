-- name: FetchUserByUsername :one
SELECT u.*
  FROM users u
  WHERE u.username = @email::text LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  name, username, encrypted_password
) VALUES (
  @name::text, @username::text, @encrypted_password::text
) RETURNING *;
