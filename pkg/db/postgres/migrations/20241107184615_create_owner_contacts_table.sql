-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS owner_contacts (
                                              owner_id INT NOT NULL REFERENCES owners(id) ON DELETE CASCADE,
    contact_id INT NOT NULL REFERENCES contacts(id) ON DELETE CASCADE,
    PRIMARY KEY (owner_id, contact_id)
    );
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO owner_contacts
VALUES
    (1,1)
    ON CONFLICT DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS owner_contacts;
-- +goose StatementEnd
