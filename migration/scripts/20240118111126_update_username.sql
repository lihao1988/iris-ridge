-- +goose Up
INSERT INTO `user` (`name`, `context`, `created_at`, `updated_at`)
VALUES ('root001', 'test_root', '2024-10-08 16:16:16', '2024-10-08 16:16:16');

-- +goose Down
DELETE FROM `user` WHERE `name` = 'root001';