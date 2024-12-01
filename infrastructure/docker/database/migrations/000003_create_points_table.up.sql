CREATE TABLE points
(
    id              uuid PRIMARY KEY NOT NULL,
    name            VARCHAR(255)     NOT NULL,
    organization_id UUID             NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP WITH TIME ZONE DEFAULT NULL,

    CONSTRAINT fk_organizations_organization FOREIGN KEY (organization_id) REFERENCES organizations (id) ON DELETE RESTRICT
);

CREATE INDEX organizations_organization_id_index ON points (organization_id);