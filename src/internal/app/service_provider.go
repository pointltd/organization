package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pointltd/organization/internal/domain/repository"
	userRepository "github.com/pointltd/organization/internal/domain/repository/user"
	"github.com/pointltd/organization/internal/infrastructure/controller"
	userController "github.com/pointltd/organization/internal/infrastructure/controller/user"
	"github.com/pointltd/organization/internal/usecase"
	createUserUseCase "github.com/pointltd/organization/internal/usecase/user"
)

type serviceProvider struct {
	db *pgxpool.Pool

	userRepository repository.UserRepository

	createUserUseCase usecase.CreateUserUseCase
	listUsersUseCase  usecase.ListUsersUseCase

	controller controller.UserController
}

func newServiceProvider(db *pgxpool.Pool) *serviceProvider {
	return &serviceProvider{
		db: db,
	}
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.db)
	}

	return s.userRepository
}

func (s *serviceProvider) CreateUserUseCase() usecase.CreateUserUseCase {
	if s.createUserUseCase == nil {
		s.createUserUseCase = createUserUseCase.NewUseCase(s.UserRepository())
	}

	return s.createUserUseCase
}

func (s *serviceProvider) ListUsersUseCase() usecase.ListUsersUseCase {
	if s.listUsersUseCase == nil {
		s.listUsersUseCase = createUserUseCase.NewListUsersUseCase(s.UserRepository())
	}

	return s.listUsersUseCase
}

func (s *serviceProvider) UserController() controller.UserController {
	if s.controller == nil {
		s.controller = userController.NewController(s.CreateUserUseCase(), s.ListUsersUseCase())
	}

	return s.controller
}
