-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS gemes (
                                     id SERIAL PRIMARY KEY,
                                     about_game TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO gemes (about_game)
VALUES ('Это начальная версия игры дающей доступ к резюме разработчика - Дениса Ханова');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS gemes;
-- +goose StatementEnd
