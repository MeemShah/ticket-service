-- +migrate Up
CREATE TABLE tickets (
    id SERIAL PRIMARY KEY,
    category VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    purchased_by VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL, 
    is_active BOOLEAN DEFAULT TRUE
);