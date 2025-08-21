CREATE TABLE IF NOT EXISTS tasks_statuses (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL,
    enabled BOOLEAN DEFAULT FALSE
);

INSERT INTO tasks_statuses (name, enabled) VALUES
    ('Logged', true),
    ('To-do', true),
    ('Doing', true),
    ('Done', true),
    ('Canceled', true);