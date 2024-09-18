CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    phone VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    name VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);