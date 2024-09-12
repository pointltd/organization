package app

import (
	"github.com/pointltd/organization/internal/domain/repository"
	userRepository "github.com/pointltd/organization/internal/domain/repository/user"
	"github.com/pointltd/organization/internal/infrastructure/controller"
	userController "github.com/pointltd/organization/internal/infrastructure/controller/user"
	"github.com/pointltd/organization/internal/usecase"
	createUserUseCase "github.com/pointltd/organization/internal/usecase/user"
)

type serviceProvider struct {
	userRepository repository.UserRepository

	createUserUseCase usecase.CreateUserUseCase

	controller controller.UserController
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

func (s *serviceProvider) UserController() controller.UserController {
	if s.controller == nil {
		s.controller = userController.NewController(s.CreateUserUseCase())
	}

	return s.controller
}
