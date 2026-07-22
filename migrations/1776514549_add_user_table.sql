-- +migrate Up
CREATE TABLE users(
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar(191) NOT NULL,
    phone_number varchar(255) NOT NULL unique,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE users;
