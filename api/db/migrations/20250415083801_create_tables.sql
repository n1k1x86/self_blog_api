-- +goose Up
-- +goose StatementBegin
CREATE TABLE tags(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE articles(
    id BIGSERIAL PRIMARY KEY,
    title TEXT,
    article_text TEXT,
    tag_id INTEGER REFERENCES tags(id),
    published_at TIMESTAMP DEFAULT NOW()

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tags;
DROP TABLE articles;
-- +goose StatementEnd
