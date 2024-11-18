-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS contacts (
                          id SERIAL PRIMARY KEY,
                          phone_number VARCHAR(15),
                          email VARCHAR(255),
                          telegram VARCHAR(50),
                          linkedin VARCHAR(255),
                          github VARCHAR(255)
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO contacts (phone_number, email, telegram, linkedin, github)
VALUES
    ('89685272694','denis_hanov@hotmail.com','@DenKhan','https://linkedin.com/in/denkhan/','https://github.com/DenisKhanov');
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS contacts;
-- +goose StatementEnd
