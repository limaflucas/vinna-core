CREATE TABLE IF NOT EXISTS habits_statuses (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR NOT NULL,
    enabled BOOLEAN DEFAULT FALSE
);

INSERT INTO habits_statuses (name, enabled) VALUES
    ('Active', TRUE),
    ('Paused', TRUE),
    ('Inactive', TRUE),
    ('Accomplished', TRUE),
    ('Abandoned', TRUE)