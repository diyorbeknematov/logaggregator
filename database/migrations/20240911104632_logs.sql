-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS logs (
    timestamp TIMESTAMP NOT NULL,
    user_id VARCHAR(100) DEFAULT '',
    service VARCHAR(100) NOT NULL,
    level VARCHAR(50) NOT NULL CHECK (level IN ('debug', 'info', 'warn', 'error', 'fatal')),
    message TEXT NOT NULL,
    error TEXT DEFAULT ''
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS logs;
-- +goose StatementEnd
