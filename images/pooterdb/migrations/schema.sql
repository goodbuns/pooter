CREATE TABLE users (
  id       SERIAL PRIMARY KEY,
  username TEXT   NOT NULL UNIQUE,
  password TEXT   NOT NULL
);

CREATE TABLE posts (
  id      SERIAL         PRIMARY KEY,
  content TEXT           NOT NULL,
  author  INTEGER        NOT NULL REFERENCES users(id),
  created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE followers (
  id       SERIAL  PRIMARY KEY,
  user_id  INTEGER NOT NULL REFERENCES users(id),
  idol     INTEGER NOT NULL REFERENCES users(id),
  CHECK (user_id != idol)
);
