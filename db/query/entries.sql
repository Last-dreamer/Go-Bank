-- name: CreateEntries :one
INSERT INTO entries (
    account_id,
    money
) VALUES (
    $1, $2
) RETURNING *;


-- name: GetEntry :one
SELECT * FROM entries
Where id = $1 LIMIT $1;

-- name: GetEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntry :one
UPDATE entries
SET money = $2
Where id = $1
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entries
Where id = $1;
