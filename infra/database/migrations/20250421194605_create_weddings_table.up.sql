-- +migrate Up
-- Create migration 20250421194605_create_weddings_table.up.sql

CREATE TABLE IF NOT EXISTS weddings (
    id            UUID            PRIMARY KEY DEFAULT gen_random_uuid(),
    date          DATE            NOT NULL,
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ     NOT NULL DEFAULT now()
);