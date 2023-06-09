CREATE TABLE IF NOT EXISTS prototypes (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT NOT NULL,
    uuid VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    short_description TEXT,
    full_description TEXT,

    engine ENUM('CONTAINER', 'VIRTUAL_MACHINE'),
    engine_ref VARCHAR(255) UNIQUE NOT NULL,

    env VARCHAR(255) NULL,
    cmd_tailer VARCHAR(255) NULL DEFAULT 'tail -f /dev/null',
    cmd VARCHAR(255) NULL DEFAULT '/bin/sh',

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)