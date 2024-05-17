-- +migrate Up
CREATE TABLE
    users (
        id BIGSERIAL PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        encrypted_password TEXT NOT NULL,
        role TEXT NOT NULL,
        created_at created_at,
        updated_at updated_at,
        deleted_at deleted_at
    );

CREATE INDEX idx_users_deleted_at ON users (deleted_at);

-- +migrate Down
DROP TABLE users;