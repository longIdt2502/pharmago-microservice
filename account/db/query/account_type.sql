-- name: GetAccountType :one
SELECT * from account_type
WHERE id = sqlc.narg(id)
    OR code = sqlc.narg(code)
LIMIT 1;

