-- name: AddRelationship :execresult
INSERT INTO relationships (UserIdFrom, UserIdTo, Relationship)
VALUES 
(?, ?, ?);

-- name: RemoveRelationship :execresult
DELETE FROM relationships WHERE UserIdFrom=? AND UserIdTo=? AND Relationship=?;

-- name: GetAllRelationshipsBetween2Users :many
SELECT * FROM relationships WHERE UserIdFrom=? AND UserIdTo=?;

-- name: GetAllUserTo :many
SELECT * FROM relationships WHERE UserIdFrom=? AND Relationship=?;

-- name: GetAllUserFrom :many
SELECT * FROM relationships WHERE UserIdTo=? AND Relationship=?;
