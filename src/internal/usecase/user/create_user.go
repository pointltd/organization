package user

import (
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/domain/repository"
	def "github.com/pointltd/organization/internal/usecase"
)

var _ def.CreateUserUseCase = (*useCase)(nil)

type CreateUserDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type useCase struct {
	userRepository repository.UserRepository
}

func NewUseCase(userRepository repository.UserRepository) *useCase {
	return &useCase{
		userRepository: userRepository,
	}
}

func (u useCase) Execute(userUUID string, info *entity.UserInfo) error {
	user := entity.User{
		UUID: userUUID,
		Info: *info,
	}

	return u.userRepository.Save(user)
}
