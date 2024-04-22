-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Project (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);
-- Create table for User
CREATE TABLE IF NOT EXISTS "user" (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    Surname VARCHAR(255) NOT NULL,
    CONSTRAINT unique_email UNIQUE (Email)
);
-- Create table for Status
CREATE TABLE IF NOT EXISTS Status (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);
-- Create table for Task
CREATE TABLE IF NOT EXISTS Task (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    StatusId INT REFERENCES Status(ID),
    ProjectId INT REFERENCES Project(ID),
    AssignedForId INT REFERENCES "user"(ID),
    DateCreated TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Project CASCADE;
DROP TABLE IF EXISTS Task CASCADE;
DROP TABLE IF EXISTS "user" CASCADE;
DROP TABLE IF EXISTS Status CASCADE;
-- +goose StatementEnd