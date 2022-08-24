ALTER TABLE users
ADD COLUMN create_at timestamp NOT NULL DEFAULT current_timestamp AFTER password
ADD COLUMN update_at timestamp NOT NULL DEFAULT current_timestamp AFTER create_at;

ALTER TABLE users ADD CONSTRAINT FK2 UNIQUE(username)
