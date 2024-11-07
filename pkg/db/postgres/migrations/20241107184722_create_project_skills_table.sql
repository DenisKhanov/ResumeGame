-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS project_skills (
                                              project_id INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    skill_id INT NOT NULL REFERENCES skills(id) ON DELETE CASCADE,
    PRIMARY KEY (project_id, skill_id)
    );
-- +goose StatementEnd


-- +goose StatementBegin
INSERT INTO project_skills
VALUES
    (1,1),
    (1,3),
    (2,1),
    (2,2),
    (2,4)
    ON CONFLICT DO NOTHING;
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS project_skills;
-- +goose StatementEnd
