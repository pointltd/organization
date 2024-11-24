package mapper

import (
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

type UserMapper interface {
	MapModelToEntity(model model.User) entity.User
	MapEntityToArg(user entity.User) pgx.NamedArgs
}
