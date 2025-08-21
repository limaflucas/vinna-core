CREATE TABLE IF NOT EXISTS habits (
    id UUID PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    name VARCHAR NOT NULL,
    start_timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    end_timestamp TIMESTAMP WITH TIME ZONE,
    frequency_id UUID REFERENCES frequencies(id),
    status INT NOT NULL REFERENCES habits_statuses(id)
);