-- +migrate Up
-- Create migration 20250421194813_create_phases_table.up.sql

CREATE TABLE IF NOT EXISTS phases (
    id                    UUID            PRIMARY KEY DEFAULT gen_random_uuid(),
    wedding_id            UUID            NOT NULL,
    months_before_wedding int             NOT NULL,
    start_date            DATE            NOT NULL,
    end_date              DATE            NOT NULL,
    FOREIGN KEY (wedding_id) REFERENCES weddings(id) ON DELETE CASCADE
);