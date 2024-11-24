package usecase

import (
	"github.com/pointltd/organization/internal/data"
	"github.com/pointltd/organization/internal/domain/entity"
)

type CreateUserUseCase interface {
	Execute(dto data.CreateUserDTO) (entity.User, error)
}

type ListUsersUseCase interface {
	Execute() ([]entity.User, error)
}
