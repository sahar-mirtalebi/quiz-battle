-- +migrate Up
INSERT IGNORE INTO permissions (title)
VALUES ('user-list'), ('user-delete');

INSERT IGNORE INTO access_controls (actor_id, actor_type, permission_id)
SELECT 2, 'role', id
FROM permissions
WHERE title IN ('user-list', 'user-delete');

-- +migrate Down
DELETE ac
FROM access_controls ac
JOIN permissions p ON p.id = ac.permission_id
WHERE ac.actor_id = 2
  AND ac.actor_type = 'role'
  AND p.title IN ('user-list', 'user-delete');

DELETE FROM permissions
WHERE title = 'user-list';
