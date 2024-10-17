-- +goose Up
CREATE TABLE demo2 (
   id SERIAL PRIMARY KEY,
   name VARCHAR(255),
   email VARCHAR(255)
);

-- +goose Down
DROP TABLE demo2;