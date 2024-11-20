package repository

import "github.com/pointltd/organization/internal/domain/entity"

type UserRepository interface {
	GetAll() ([]entity.User, error)
	Save(user entity.User) error
}
