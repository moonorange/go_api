-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `todos` (
    `id` CHAR(36) NOT NULL PRIMARY KEY,
    `description` VARCHAR(255) COLLATE utf8mb4_bin NOT NULL,
    `is_completed` BOOLEAN NOT NULL DEFAULT FALSE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `todos`;
-- +goose StatementEnd
