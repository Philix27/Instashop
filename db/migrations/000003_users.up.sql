CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  wallets TEXT[],
  first_name TEXT,
  last_name TEXT,
  middle_name TEXT,
  dob TIMESTAMP,
  email TEXT  UNIQUE  NOT NULL,
  phone TEXT  UNIQUE,
  hashed_password TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
