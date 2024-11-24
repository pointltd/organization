package model

import (
	"database/sql"
)

type User struct {
	ID          string         `db:"id"`
	FirstName   string         `db:"first_name"`
	LastName    sql.NullString `db:"last_name"`
	Email       string         `db:"email"`
	Password    string         `db:"password"`
	Phone       sql.NullString `db:"phone"`
	CreatedAt   sql.NullTime   `db:"created_at"`
	UpdatedAt   sql.NullTime   `db:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at"`
	CreatedById sql.NullString `db:"created_by_id"`
	UpdatedById sql.NullString `db:"updated_by_id"`
	DeletedById sql.NullString `db:"deleted_by_id"`
}
