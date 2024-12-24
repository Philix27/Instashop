CREATE TABLE products (
  id   BIGSERIAL PRIMARY KEY,
  title TEXT   UNIQUE     NOT NULL,
  description TEXT   UNIQUE     NOT NULL,
  image_url TEXT   UNIQUE     NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

