-- name: CreateUser :one
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING id, name, dob;

-- name: GetUser :one
SELECT id, name, 
       TO_CHAR(dob, 'YYYY-MM-DD')::text AS dob, 
       EXTRACT(YEAR FROM AGE(dob))::int AS age
FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET name = $2,
    dob = $3
WHERE id = $1
RETURNING id, name, TO_CHAR(dob, 'YYYY-MM-DD')::text AS dob;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, 
       TO_CHAR(dob, 'YYYY-MM-DD')::text AS dob, 
       EXTRACT(YEAR FROM AGE(dob))::int AS age
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;