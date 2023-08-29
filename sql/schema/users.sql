DROP TABLE IF EXISTS tokens;

DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS accounts;

CREATE TABLE
    users (
        id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
        email varchar NOT NULL UNIQUE,
        name TEXT NOT NULL,
        password TEXT NOT NULL,
        last_accessed TIMESTAMP NOT NULL DEFAULT current_timestamp,
        created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
        updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
    );

CREATE TABLE
    tokens (
        id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
        user_id UUID NOT NULL REFERENCES users(id),
        session_token VARCHAR(255) NOT NULL UNIQUE,
        created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
        last_accessed_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
        expires_at TIMESTAMP NOT NULL DEFAULT current_timestamp + interval '7 DAY',
        is_valid BOOLEAN NOT NULL DEFAULT true
    );

CREATE TABLE accounts (
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    user_name VARCHAR(255) NOT NULL UNIQUE,
    picture VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);