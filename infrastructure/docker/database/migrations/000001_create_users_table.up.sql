SET TIME ZONE 'UTC';

CREATE TABLE users
(
    id            uuid PRIMARY KEY    NOT NULL,
    first_name    VARCHAR(255)        NOT NULL,
    last_name     VARCHAR(255)        NULL,
    email         VARCHAR(255) UNIQUE NOT NULL,
    password      VARCHAR(255)        NOT NULL,
    phone         VARCHAR(20)         NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP WITH TIME ZONE DEFAULT NULL
);