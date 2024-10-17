-- +goose Up
CREATE TABLE demo (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255),
   email VARCHAR(255)
);

-- +goose Down
DROP TABLE demo;