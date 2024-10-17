CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL,
  username TEXT NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  PRIMARY KEY (id)
);
