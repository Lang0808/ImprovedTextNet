-- name: CreateBlog :execresult
INSERT INTO blogs (Title, Privacy,CreatedTime, UpdatedTime,Author, Content)
VALUES
(?, ?, ?, ?, ?, ?);

-- name: GetBlog :one
SELECT * FROM blogs
WHERE Id=?;

-- name: GetBlogsOfAuthor :many
SELECT * FROM blogs
WHERE Author=? 
ORDER BY Id DESC
LIMIT ? OFFSET ?;