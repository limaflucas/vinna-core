CREATE TABLE IF NOT EXISTS users_statuses (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL,
    enabled BOOLEAN DEFAULT FALSE
);

INSERT INTO tasks_statuses (name, enabled) VALUES
    ('Active', true),
    ('Blocked', true),
    ('Deleted', true);