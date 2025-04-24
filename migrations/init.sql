CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    item VARCHAR(255),
    price float
);