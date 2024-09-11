package user

import (
	"github.com/pointltd/organization/internal/domain/entity"
	def "github.com/pointltd/organization/internal/domain/repository"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) Save(user entity.User) error {
	return nil
}
