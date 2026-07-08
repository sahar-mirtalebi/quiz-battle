-- +migrate Up
ALTER TABLE users ADD COLUMN password varchar(191) NOT NULL;

-- +migrate Down
ALTER TABLE users DROP COLUMN password;