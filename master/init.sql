CREATE DATABASE demo;

\c demo

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT,
  email TEXT
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_name ON users(name);

INSERT INTO users (name, email)
SELECT
  md5(random()::text),
  md5(random()::text) || '@mail.com'
FROM generate_series(1, 1000000);
