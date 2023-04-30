CREATE TABLE relationships (
    UserIdFrom INTEGER NOT NULL,
    UserIdTo INTEGER NOT NULL,
    Relationship INTEGER NOT NULL,
    PRIMARY KEY(UserIdFrom, UserIdTo, Relationship),
    INDEX(UserIdFrom, Relationship),
    INDEX(UserIdTo, Relationship)
);