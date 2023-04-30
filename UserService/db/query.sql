-- name: CreateUser :execresult
INSERT INTO users (Username, Avatar, Password)
VALUES
(?, ?, ?);

-- name: GetUserFromUsername :one
SELECT * FROM users WHERE Username=?;

-- name: GetUserFromUserId :one
SELECT * FROM users WHERE Id=?;
