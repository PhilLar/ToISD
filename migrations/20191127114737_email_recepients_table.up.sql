CREATE TABLE IF NOT EXISTS email_recepients
(
  id            SERIAL PRIMARY KEY,
  "name"        TEXT NOT NULL,
  email         TEXT NOT NULL
);