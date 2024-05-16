
-- +migrate Up
CREATE TABLE
    chats (
        id BIGSERIAL PRIMARY KEY,
        telegram_chat_id BIGINT UNIQUE,
        created_at created_at,
        updated_at updated_at,
        deleted_at deleted_at
    );

CREATE INDEX idx_chats_deleted_at ON chats (deleted_at);

-- +migrate Down
DROP TABLE chats;