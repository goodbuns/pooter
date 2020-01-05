CREATE TABLE users (
  id       SERIAL PRIMARY KEY,
  username TEXT   NOT NULL UNIQUE,
  password TEXT   NOT NULL
);

SET TIME ZONE "PST8PDT";

CREATE TABLE posts (
  id      SERIAL PRIMARY KEY,
  content TEXT,
  user_id  INTEGER REFERENCES users(id),
  created_at TIMESTAMPTZ
);

CREATE TABLE followers (
  id          SERIAL  PRIMARY KEY,
  user_id      INTEGER REFERENCES users(id),
  follow_id    INTEGER REFERENCES users(id),
  CHECK (user_id != follow_id)
);
