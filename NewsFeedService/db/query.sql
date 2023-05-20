-- name: AddNewsFeed :execresult
INSERT INTO NewsFeed (UserId, BlogId, Score)
VALUES
(?, ?, ?);

-- name: GetNewsFeed :many
SELECT * FROM NewsFeed
WHERE UserId=?
ORDER BY Score DESC
LIMIT ? OFFSET ?;