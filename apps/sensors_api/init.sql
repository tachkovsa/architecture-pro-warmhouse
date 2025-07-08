-- Connect to the database
\c smarthome;

CREATE TABLE IF NOT EXISTS sensors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(50) NOT NULL,
    location VARCHAR(100) NOT NULL,
    value FLOAT DEFAULT 0,
    unit VARCHAR(20),
    status VARCHAR(20) NOT NULL DEFAULT 'inactive',
    last_updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
); 