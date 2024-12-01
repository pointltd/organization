package model

import "database/sql"

type Organization struct {
	Id        string       `db:"id"`
	Name      string       `db:"name"`
	OwnerId   string       `db:"owner_id"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
