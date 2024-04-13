
-- +goose Up
CREATE TABLE banners (
    feature INT,
    tag INT,
    json_content JSONB NOT NULL,
    is_active BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    PRIMARY KEY (feature, tag)
);

-- +goose Down
DROP TABLE banners;
