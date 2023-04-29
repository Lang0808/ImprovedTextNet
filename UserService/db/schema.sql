CREATE TABLE users (
    Id INTEGER AUTO_INCREMENT,
    Username NVARCHAR(50) NOT NULL,
    Avatar NVARCHAR(50),
    Password NVARCHAR(200),
    PRIMARY KEY(id),
    UNIQUE(Username)
);