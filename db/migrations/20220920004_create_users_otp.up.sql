CREATE TABLE users_otp
(
    user_id int(10) NOT NULL,
    otp int(6) NOT NULL,
    create_at timestamp NOT NULL DEFAULT current_timestamp,
    expire_at timestamp NOT NULL DEFAULT current_timestamp,
     FOREIGN KEY (user_id) REFERENCES users (id)
);