CREATE TABLE exchange_rates (
    id SERIAL PRIMARY KEY,
    base_currency INT NOT NULL REFERENCES currencies(id),
    target_currency INT NOT NULL REFERENCES currencies(id),
    rate NUMERIC(12, 6) NOT NULL,
    updated_at TIMESTAMP DEFAULT NOW()
);

