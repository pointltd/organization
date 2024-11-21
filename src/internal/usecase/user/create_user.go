package user

import (
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/domain/repository"
	def "github.com/pointltd/organization/internal/usecase"
)

var _ def.CreateUserUseCase = (*createUserUseCase)(nil)

type CreateUserDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type createUserUseCase struct {
	userRepository repository.UserRepository
}

func NewUseCase(userRepository repository.UserRepository) *createUserUseCase {
	return &createUserUseCase{
		userRepository: userRepository,
	}
}

func (u createUserUseCase) Execute(userID string, info *entity.UserInfo) error {
	user := entity.User{
		ID:   userID,
		Info: *info,
	}

	return u.userRepository.Save(user)
}
