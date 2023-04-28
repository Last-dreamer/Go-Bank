-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    money
) VALUES(
  $1, $2, $3
) RETURNING *;


-- name: GetTransfer :one
SELECT * FROM transfers
Where id = $1 LIMIT $1;


-- name: GetListTransfers :many
SELECT * from transfers
Where id = $1
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers
SET money  = $2
Where id = $1
RETURNING *;


-- name: DeleteTransfer :exec
DELETE FROM transfers
Where id = $1;