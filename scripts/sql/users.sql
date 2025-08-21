CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL,
    status_id INT NOT NULL REFERENCES users_statuses(id),
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    country VARCHAR NOT NULL,
    birth DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);