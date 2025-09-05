CREATE TABLE IF NOT EXISTS orders (
    order_id VARCHAR(100) PRIMARY KEY CHECK (order_id <> ''),
    user_id VARCHAR(100) NOT NULL CHECK (user_id <> ''),
    created_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_deleted_at ON orders(deleted_at);