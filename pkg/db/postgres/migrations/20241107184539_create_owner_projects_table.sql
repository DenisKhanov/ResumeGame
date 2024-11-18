-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS owner_projects (
                                              owner_id INT NOT NULL REFERENCES owners(id) ON DELETE CASCADE,
    project_id INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    PRIMARY KEY (owner_id, project_id)
    );
-- +goose StatementEnd


-- +goose StatementBegin
INSERT INTO owner_projects
VALUES
    (1,1),
    (1,2)
    ON CONFLICT DO NOTHING;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS owner_projects;
-- +goose StatementEnd
