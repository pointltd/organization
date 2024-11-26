package auth

import (
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/domain/repository"
	def "github.com/pointltd/organization/internal/usecase"
	passwordPkg "github.com/pointltd/organization/pkg/password"
)

var _ def.AuthenticateUserUseCase = (*authenticateUserUseCase)(nil)

type authenticateUserUseCase struct {
	userRepository repository.UserRepository
}

func NewAuthenticateUserUseCase(userRepository repository.UserRepository) *authenticateUserUseCase {
	return &authenticateUserUseCase{
		userRepository: userRepository,
	}
}

func (u authenticateUserUseCase) Execute(email string, password string) (*entity.User, error) {
	user, err := u.userRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	if !passwordPkg.VerifyPassword(password, user.Password) {
		return nil, nil
	}

	return user, nil
}
