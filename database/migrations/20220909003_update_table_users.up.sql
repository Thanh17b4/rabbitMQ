ALTER TABLE users
ADD COLUMN created_at timestamp NOT NULL DEFAULT current_timestamp AFTER password;

