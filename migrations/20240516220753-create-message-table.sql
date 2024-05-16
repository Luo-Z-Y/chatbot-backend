-- +migrate Up
CREATE TABLE
    messages (
        id BIGSERIAL PRIMARY KEY,
        telegram_message_id BIGINT NOT NULL,
        by TEXT NOT NULL,
        message_body TEXT NOT NULL ,
        "timestamp" timestamptz NOT NULL,
        hotel_staff_id BIGINT REFERENCES users (id),
        request_query_id BIGINT NOT NULL REFERENCES request_queries (id),
        created_at created_at,
        updated_at updated_at,
        deleted_at deleted_at
    );

CREATE INDEX idx_messages_deleted_at ON messages (deleted_at);

-- +migrate Down
DROP TABLE messages;