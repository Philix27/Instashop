CREATE TABLE orders (
  id   BIGSERIAL PRIMARY KEY,
  user_id  INT REFERENCES users (id),
  order_status OrderStatus,
  items orderItems[],
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


