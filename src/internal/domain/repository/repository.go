package repository

import "github.com/pointltd/organization/internal/domain/entity"

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Save(user entity.User) (entity.User, error)
	GetOrganizations(userId string) ([]entity.Organization, error)
}

type OrganizationRepository interface {
	Save(organization entity.Organization) (entity.Organization, error)
}

type PointRepository interface {
	Save(point entity.Point) (entity.Point, error)
}
