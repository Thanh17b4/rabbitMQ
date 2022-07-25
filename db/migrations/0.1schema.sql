CREATE TABLE users (
        id        int(10) NOT NULL AUTO_INCREMENT,
        name      varchar(255) NOT NULL,
        email     varchar (255) NOT NULL,
        protected   tinyint(1) NOT NULL,
        banned    	tinyint(1) NOT NULL,
        activated 	tinyint(1) NOT NULL,
        address   varchar (255) NOT NULL,
        password  varchar(255) NOT NULL,
        username  varchar(255) NOT NULL,
                   PRIMARY KEY (id),
            CONSTRAINT users_email_unique UNIQUE(email)
)
CREATE TABLE users_otp (
    user_id int(10) NOT NULL,
    otp int(6) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expired_at timestamp NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id)
)