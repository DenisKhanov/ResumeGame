-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS skills (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(100) NOT NULL,
                        description TEXT,
                        level VARCHAR(50) NOT NULL CHECK (level IN ('Beginner', 'Intermediate', 'Advanced', 'Expert'))
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO skills (name, description, level)
VALUES
    ('Golang','Programming language','Advanced'),
    ('Redis','Key/value in memory storage','Beginner'),
    ('REST API','Representational State Transfer — architect stile API','Advanced'),
    ('gRPC','Open source high performance Remote Procedure Call (RPC) framework','Intermediate');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS skills;
-- +goose StatementEnd
