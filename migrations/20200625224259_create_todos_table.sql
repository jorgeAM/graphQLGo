-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE todos(
  id SERIAL PRIMARY KEY,
  title VARCHAR(100),
  description TEXT,
  user_id  INTEGER REFERENCES users (id),
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE todos;