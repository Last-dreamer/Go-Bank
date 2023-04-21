-- name: CreateAccont :one
INSERT INTO account (
    owner,
    balance,
    currency
) VALUES ( 
    $1, $2, $3
    ) RETURNING *;

-- name: GetAccount :one
SELECT * FROM account 
Where id = $1 LIMIT $1;

-- name: GetListAccount :many
SELECT * FROM account
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE account
SET balance = $2
Where id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM account
Where id = $1;


