-- name: GetAllUsers :many

select * from users;

-- name: GetOneUserWithEmail :one

select * from users where email = $1 LIMIT 1;

-- name: CreateUser :one

INSERT into
    users (email, name, password)
values ($1, $2, $3) RETURNING *;

-- name: CreateSession :exec

INSERT INTO tokens(user_id, session_token) VALUES($1, $2);

-- name: GetLatestSessionWithUserId :one

SELECT *
FROM tokens
WHERE user_id = $1
ORDER BY expires_at DESC
LIMIT 1;

-- name: InvalidateTokenAll :exec

UPDATE tokens SET is_valid = false WHERE user_id = $1;

-- name: InvalidateTokenLatest :exec

UPDATE tokens SET is_valid = false WHERE id = $1;