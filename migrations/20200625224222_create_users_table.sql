-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users(
  id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  email VARCHAR(100) UNIQUE,
  password VARCHAR(100),
  created_at TIMESTAMPTZ DEFAULT NOW(),
  Updated_at TIMESTAMPTZ DEFAULT NOW()
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;
