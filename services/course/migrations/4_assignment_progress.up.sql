CREATE TABLE IF NOT EXISTS assignment_progress (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    assignment_uuid VARCHAR(255) NOT NULL,
    user_uuid VARCHAR(255) NOT NULL,
    environment_uuid VARCHAR(255) NOT NULL,

    grade INT UNSIGNED DEFAULT NULL,

    FOREIGN KEY (user_uuid) REFERENCES users (uuid),
    FOREIGN KEY (environment_uuid) REFERENCES environments (uuid),

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);