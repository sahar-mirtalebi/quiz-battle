-- +migrate Up
CREATE TABLE access_controls (
    id INT AUTO_INCREMENT PRIMARY KEY,
    actor_id INT NOT NULL,
    actor_type ENUM('role', 'user') NOT NULL,
    permission_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE INDEX idx_access_controls_actor_permission (actor_id, actor_type, permission_id),
    FOREIGN KEY (permission_id) REFERENCES permissions(id)
);

-- +migrate Down
DROP TABLE access_controls;
