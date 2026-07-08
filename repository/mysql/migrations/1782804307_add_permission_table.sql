-- +migrate Up
CREATE TABLE permissions(
    id int AUTO_INCREMENT PRIMARY KEY,
    title varchar(191) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE permissions;