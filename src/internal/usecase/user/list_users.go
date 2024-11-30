package user

import (
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/domain/repository"
	def "github.com/pointltd/organization/internal/usecase"
)

var _ def.ListUsersUseCase = (*listUsersUseCase)(nil)

type listUsersUseCase struct {
	userRepository repository.UserRepository
}

func NewListUsersUseCase(userRepository repository.UserRepository) *listUsersUseCase {
	return &listUsersUseCase{
		userRepository: userRepository,
	}
}

func (u listUsersUseCase) Execute() ([]entity.User, error) {
	return u.userRepository.GetAll()
}
