CREATE TABLE IF NOT EXISTS items
(
  id        SERIAL PRIMARY KEY,
  title     TEXT NOT NULL,
  descr     TEXT NOT NULL,
  price     TEXT NOT NULL,
  amount    SMALLINT NOT NULL,
  user_id   INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE
);