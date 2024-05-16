-- +migrate Up
CREATE TABLE
    bookings (
        id BIGSERIAL PRIMARY KEY,
        room_number TEXT,
        last_name TEXT,
        chat_id BIGINT NOT NULL REFERENCES chats (id),
        created_at created_at,
        updated_at updated_at,
        deleted_at deleted_at
    );

CREATE INDEX idx_bookings_deleted_at ON bookings (deleted_at);

-- +migrate Down
DROP TABLE bookings;
