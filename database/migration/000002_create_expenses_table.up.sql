CREATE TABLE expenses (
    id SERIAL PRIMARY KEY,
    description VARCHAR(50) NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    currency INT NOT NULL REFERENCES currencies(id),
    created_at TIMESTAMP DEFAULT NOW()
);