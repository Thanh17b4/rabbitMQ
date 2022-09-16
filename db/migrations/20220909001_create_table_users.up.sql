CREATE TABLE users
(
    id        int(10) NOT NULL AUTO_INCREMENT,
    name      varchar(255) NOT NULL,
    email     varchar(255) NOT NULL,
    protected tinyint(1) NOT NULL,
    banned    tinyint(1) NOT NULL,
    activated tinyint(1) NOT NULL,
    address   varchar(255) NOT NULL,
    password  varchar(255) NOT NULL,
    username  varchar(255) NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT users_email_unique UNIQUE (email)
);