CREATE TABLE IF NOT EXISTS user_cart_items (
    product_id VARCHAR (100) NOT NULL CHECK (product_id <> ''),
    user_id VARCHAR (100) NOT NULL CHECK (user_id <> ''),
    name VARCHAR (255) NOT NULL CHECK (name <> ''),
    description TEXT,
    price REAL NOT NULL CHECK (price > 0),
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    PRIMARY KEY (product_id, user_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_user_cart_items_price ON user_cart_items(price);
CREATE INDEX IF NOT EXISTS idx_user_cart_items_user ON user_cart_items(user_id);