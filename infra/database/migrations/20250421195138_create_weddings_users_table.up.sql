-- +migrate Up
-- Create migration 20250421195138_create_weddings_users_table.up.sql

CREATE TABLE IF NOT EXISTS wedding_users (
    wedding_id    UUID            NOT NULL,
    user_id       UUID            NOT NULL,
    PRIMARY KEY (wedding_id, user_id),
    FOREIGN KEY (wedding_id) REFERENCES weddings(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id)    REFERENCES users(id)    ON DELETE CASCADE
);