-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS owner_experience (
                                                owner_id INT NOT NULL REFERENCES owners(id) ON DELETE CASCADE,
    experience_id INT NOT NULL REFERENCES experience(id) ON DELETE CASCADE,
    PRIMARY KEY (owner_id, experience_id)
    );
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO owner_experience
VALUES
    (1,1),
    (1,2),
    (1,3),
    (1,4)
    ON CONFLICT DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS owner_experience;
-- +goose StatementEnd
