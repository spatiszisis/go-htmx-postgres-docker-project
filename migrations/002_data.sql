-- +goose Up
-- +goose StatementBegin
-- Status ENUM('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') NOT NULL DEFAULT 'TODO',
INSERT INTO status
("name")
VALUES('TODO');
INSERT INTO status
("name")
VALUES('DONE');

INSERT INTO "user"
( "name", surname, email)
VALUES('John', 'Doe', 'john@example.com');

INSERT INTO project
("name")
VALUES('Project1');
-- +goose StatementEnd

-- +goose Down