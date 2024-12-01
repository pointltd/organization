package usecase

import (
	"github.com/pointltd/organization/internal/data"
	"github.com/pointltd/organization/internal/domain/entity"
)

type AuthenticateUserUseCase interface {
	Execute(email string, password string) (*entity.User, error)
}

type CreateUserUseCase interface {
	Execute(dto data.CreateUserDTO) (entity.User, error)
}

type ListUsersUseCase interface {
	Execute() ([]entity.User, error)
}

type ListUserOrganizationsUseCase interface {
	Execute(userId string) ([]entity.Organization, error)
}

type CreateOrganizationUseCase interface {
	Execute(name string, ownerId string) (entity.Organization, error)
}

type CreatePointUseCase interface {
	Execute(name string, organizationId string) (entity.Point, error)
}

type ListOrganizationPointsUseCase interface {
	Execute(organizationId string) ([]entity.Point, error)
}
