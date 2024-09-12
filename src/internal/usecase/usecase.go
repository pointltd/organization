package usecase

import "github.com/pointltd/organization/internal/domain/entity"

type CreateUserUseCase interface {
	Execute(userUUID string, info *entity.UserInfo) error
}
