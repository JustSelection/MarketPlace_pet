CREATE TABLE IF NOT EXISTS order_items (
                                           order_id VARCHAR(100) NOT NULL CHECK (order_id <> ''),
    product_id VARCHAR(100)  NOT NULL CHECK (product_id <> ''),
    user_id VARCHAR(100) NOT NULL CHECK (user_id <> ''),
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    price REAL NOT NULL CHECK (price > 0),
    deleted_at TIMESTAMP WITH TIME ZONE,
                             PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
    );

CREATE INDEX IF NOT EXISTS idx_order_items_product_id ON order_items(product_id);
CREATE INDEX IF NOT EXISTS idx_order_items_user_id ON order_items(user_id);
CREATE INDEX IF NOT EXISTS idx_order_items_deleted_at ON order_items(deleted_at);