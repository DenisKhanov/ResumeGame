-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS owner_skills (
                                            owner_id INT NOT NULL REFERENCES owners(id) ON DELETE CASCADE,
    skill_id INT NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    PRIMARY KEY (owner_id, skill_id)
    );
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO owner_skills
VALUES
    (1,1),
    (1,2),
    (1,3),
    (1,4)
    ON CONFLICT DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS owner_skills;
-- +goose StatementEnd
