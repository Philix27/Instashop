CREATE TABLE order_items (
  id   BIGSERIAL PRIMARY KEY,
  quantity INT NOT NULL CHECK (quantity > 0),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   order_id  INT REFERENCES orders (id),
   product_id  INT REFERENCES products (id)
);

