-- +goose Up
-- +goose StatementBegin
CREATE TABLE photos(
    id UUID PRIMARY KEY, -- also used as the base object name in storage
    alt_text TEXT,
    caption TEXT,
    uploaded_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE photos;
-- +goose StatementEnd
