package mapper

import (
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

type UserMapper interface {
	MapModelToEntity(model model.User) (entity.User, error)
}
