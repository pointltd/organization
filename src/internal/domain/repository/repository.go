package repository

import "github.com/pointltd/organization/internal/domain/entity"

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Save(user entity.User) (entity.User, error)
}
