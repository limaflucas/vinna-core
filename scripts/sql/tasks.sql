CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    name VARCHAR NOT NULL,
    start_timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    end_timestamp TIMESTAMP WITH TIME ZONE,
    frequency_id UUID REFERENCES frequencies(id),
    status_id INT NOT NULL REFERENCES tasks_statuses(id)
);