CREATE TABLE users (
  id       SERIAL PRIMARY KEY,
  username TEXT   NOT NULL UNIQUE,
  password TEXT   NOT NULL
);

CREATE TABLE posts (
  id      SERIAL PRIMARY KEY,
  content TEXT,
  userid  INTEGER REFERENCES users(id)
);

CREATE TABLE followers (
  id          SERIAL  PRIMARY KEY,
  userid      INTEGER REFERENCES users(id),
  followid    INTEGER REFERENCES users(id),
  CHECK (userid != followid)
);
