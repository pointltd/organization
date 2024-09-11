package repository

import "github.com/pointltd/organization/internal/domain/entity"

type UserRepository interface {
	Save(user entity.User) error
}
