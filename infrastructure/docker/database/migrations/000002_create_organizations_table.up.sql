CREATE TABLE organizations
(
    id            uuid PRIMARY KEY         DEFAULT gen_random_uuid(),
    name          VARCHAR(255) NOT NULL,
    owner_id      UUID         NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    created_by_id UUID         NULL,
    updated_by_id UUID         NULL,
    deleted_by_id UUID         NULL,

    CONSTRAINT fk_users_owner FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE RESTRICT,
    CONSTRAINT fk_users_created_by FOREIGN KEY (created_by_id) REFERENCES users (id) ON DELETE SET NULL,
    CONSTRAINT fk_users_updated_by FOREIGN KEY (updated_by_id) REFERENCES users (id) ON DELETE SET NULL,
    CONSTRAINT fk_users_deleted_by FOREIGN KEY (deleted_by_id) REFERENCES users (id) ON DELETE SET NULL
);