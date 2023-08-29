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

-- name: GetAccountByUserId :one

SELECT * FROM tokens WHERE user_id = $1 LIMIT 1;

-- name: CreateAccount :exec

INSERT INTO accounts (user_id, user_name) VALUES ($1, $2);

-- name: UpdateAccount :one

UPDATE accounts AS a
SET
    user_name = $2,
    picture = $3,
    updated_at = current_timestamp
FROM (
        VALUES (
                $1,
                'new_user_name',
                'new_picture'
            )
    ) AS upd(id, user_name, picture)
WHERE
    a.id = upd.id
    AND (
        a.user_name IS DISTINCT
        FROM
            upd.user_name
            OR a.picture IS DISTINCT
        FROM
            upd.picture
    )
    AND a.id = $1 RETURNING *;