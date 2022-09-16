ALTER TABLE users
ADD COLUMN create_at timestamp NOT NULL DEFAULT current_timestamp AFTER password;

ALTER TABLE users ADD CONSTRAINT users_user_name_unique UNIQUE(username);

ALTER TABLE users
ADD COLUMN update_at timestamp NOT NULL DEFAULT current_timestamp AFTER create_at;