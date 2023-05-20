-- name: AddNewsFeed :execresult
INSERT INTO NewsFeed (UserId, BlogId, Score)
VALUES
(?, ?, ?);

-- name: GetNewsFeed :execresult
SELECT * FROM NewsFeed
WHERE UserId=?
ORDER BY Score DESC
LIMIT ? OFFSET ?;