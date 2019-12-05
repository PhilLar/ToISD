CREATE TABLE IF NOT EXISTS users
(
  id            SERIAL PRIMARY KEY,
  "name"        TEXT NOT NULL,
  email         TEXT NOT NULL,
  "address"     TEXT NOT NULL,
  "password"    TEXT NOT NULL,
  phone_number  TEXT NOT NULL
);