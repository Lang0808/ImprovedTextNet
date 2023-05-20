CREATE TABLE NewsFeed (
    UserId INTEGER NOT NULL,
    BlogId INTEGER NOT NULL,
    Score INTEGER NOT NULL,
    PRIMARY KEY(UserId, BlogId),
    INDEX(UserId, Score)
);

