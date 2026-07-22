-- +migrate Up
INSERT IGNORE INTO permissions (title)
VALUES ('user-delete');

-- +migrate Down
DELETE FROM permissions
WHERE title = 'user-delete';
