CREATE TABLE organizations
(
    id            uuid PRIMARY KEY NOT NULL,
    name          VARCHAR(255)     NOT NULL,
    owner_id      UUID             NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP WITH TIME ZONE DEFAULT NULL,

    CONSTRAINT fk_users_owner FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE RESTRICT
);

CREATE INDEX organizations_owner_id_index ON organizations (owner_id);