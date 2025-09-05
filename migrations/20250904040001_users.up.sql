CREATE TABLE IF NOT EXISTS users (
    user_id VARCHAR(100) PRIMARY KEY CHECK (user_id <> ''),
    name VARCHAR(255) NOT NULL CHECK(name <> ''),
    information TEXT,
    email VARCHAR(255) NOT NULL UNIQUE CHECK(email <> ''),
    password VARCHAR(255) NOT NULL CHECK (password <> ''),
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users(email) WHERE deleted_at IS NULL;