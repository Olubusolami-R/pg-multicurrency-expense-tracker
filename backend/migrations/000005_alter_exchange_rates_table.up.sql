ALTER TABLE exchange_rates
ADD CONSTRAINT exchange_rates_base_target_unique UNIQUE (base_currency, target_currency);
