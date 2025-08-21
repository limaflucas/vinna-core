CREATE TABLE IF NOT EXISTS frequencies (
    id UUID PRIMARY KEY NOT NULL,
    "minute" VARCHAR NOT NULL,
    "hour" VARCHAR NOT NULL,
    day_of_month VARCHAR NOT NULL,
    "month" VARCHAR NOT NULL,
    day_of_week VARCHAR NOT NULL
)