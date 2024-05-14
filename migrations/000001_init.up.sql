-- +migrate Up
CREATE TABLE IF NOT EXISTS valutes (
    id SERIAL PRIMARY KEY,
    curs DOUBLE PRECISION,
    vch_code VARCHAR(10),
    created_at TIMESTAMP
    );