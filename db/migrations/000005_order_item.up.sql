CREATE TABLE order_items (
  id   BIGSERIAL PRIMARY KEY,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  quantity INT NOT NULL CHECK (quantity > 0),
   order_id  INT REFERENCES orders (id),
   product_id  INT REFERENCES products (id),
   user_id  INT REFERENCES users (id)
);

