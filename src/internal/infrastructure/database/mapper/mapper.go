package mapper

import (
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
)

type UserMapper interface {
	MapRowToUser(rows pgx.Rows) (entity.User, error)
}
