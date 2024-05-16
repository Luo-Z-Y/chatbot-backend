-- +migrate Up
CREATE TABLE
    request_queries (
        id BIGSERIAL PRIMARY KEY,
        status TEXT,
        type TEXT,
        chat_id BIGINT REFERENCES chats (id),
        booking_id BIGINT REFERENCES bookings (id),
        created_at created_at,
        updated_at updated_at,
        deleted_at deleted_at
    );

CREATE INDEX idx_request_queries_deleted_at ON request_queries (deleted_at);

-- +migrate Down
DROP TABLE request_queries;
