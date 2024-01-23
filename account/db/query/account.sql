-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountByUseName :one
SELECT * FROM accounts
WHERE username = $1 LIMIT 1;

-- name: CreateAccount :one
INSERT INTO accounts (username, hashed_password, full_name, email, type)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateAccount :one
UPDATE accounts
SET
    is_verify = COALESCE(sqlc.narg(is_verify), is_verify)
WHERE
    id = sqlc.narg(id)
    OR username = sqlc.narg(username)
RETURNING *;
