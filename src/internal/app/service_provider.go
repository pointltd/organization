package app

import (
	"github.com/pointltd/organization/internal/domain/repository"
	userRepository "github.com/pointltd/organization/internal/domain/repository/user"
	"github.com/pointltd/organization/internal/usecase"
	createUserUseCase "github.com/pointltd/organization/internal/usecase/user"
)

type serviceProvider struct {
	userRepository repository.UserRepository

	createUserUseCase usecase.CreateUserUseCase
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository()
	}

	return s.userRepository
}

func (s *serviceProvider) CreateUserUseCase() usecase.CreateUserUseCase {
	if s.createUserUseCase == nil {
		s.createUserUseCase = createUserUseCase.NewUseCase(s.UserRepository())
	}

	return s.createUserUseCase
}
