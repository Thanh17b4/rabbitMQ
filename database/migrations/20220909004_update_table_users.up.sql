
ALTER TABLE users
ADD COLUMN updated_at timestamp NOT NULL DEFAULT current_timestamp AFTER created_at;