-- name: CreateVerify :one
INSERT INTO verifies (
    username, email, secret_code
) VALUES ($1, $2, $3) RETURNING *;

-- name: GetVerify :one
SELECT * FROM verifies
WHERE id = $1 LIMIT 1;

-- name: UpdateVerify :one
UPDATE verifies
SET
    is_used = TRUE
WHERE
    id = $1
RETURNING *;