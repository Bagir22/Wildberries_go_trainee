CREATE TABLE IF NOT EXISTS orders
(
    id          SERIAL PRIMARY KEY,
    data        jsonb NOT NULL
);