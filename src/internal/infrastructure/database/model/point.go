package model

import "database/sql"

type Point struct {
	Id             string       `db:"id"`
	Name           string       `db:"name"`
	OrganizationId string       `db:"organization_id"`
	CreatedAt      sql.NullTime `db:"created_at"`
	UpdatedAt      sql.NullTime `db:"updated_at"`
	DeletedAt      sql.NullTime `db:"deleted_at"`
}
