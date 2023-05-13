CREATE TABLE blogs (
    Id INTEGER AUTO_INCREMENT,
    Title NVARCHAR(100) NOT NULL,
    Privacy Integer NOT NULL,
    CreatedTime BIGINT NOT NULL,
    UpdatedTime BIGINT NOT NULL,
    Author INTEGER NOT NULL,
    Content TEXT NOT NULL,
    PRIMARY KEY(Id)
);