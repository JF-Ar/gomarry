-- +migrate Up
-- Create migration 20250421200849_create_tasks_table.up.sql
CREATE TYPE status AS ENUM ('pending', 'in_progress', 'done', 'need_help');

CREATE TABLE IF NOT EXISTS tasks (
    id            UUID            PRIMARY KEY DEFAULT gen_random_uuid(),
    phase_id      UUID            NOT NULL,
    title         TEXT            NOT NULL,
    description   TEXT            NULL,
    assigned_to   UUID            NULL,
    due_date      DATE            NOT NULL,
    status        status          DEFAULT NULL,
    FOREIGN KEY (phase_id)    REFERENCES phases(id) ON DELETE CASCADE,
    FOREIGN KEY (assigned_to) REFERENCES users(id) ON DELETE SET NULL
);