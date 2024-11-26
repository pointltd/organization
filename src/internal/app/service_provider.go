package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pointltd/organization/internal/domain/repository"
	userRepository "github.com/pointltd/organization/internal/domain/repository/user"
	"github.com/pointltd/organization/internal/infrastructure/database/mapper"
	userMapper "github.com/pointltd/organization/internal/infrastructure/database/mapper/user"
	"github.com/pointltd/organization/internal/infrastructure/http/controller"
	authController "github.com/pointltd/organization/internal/infrastructure/http/controller/auth"
	userController "github.com/pointltd/organization/internal/infrastructure/http/controller/user"
	"github.com/pointltd/organization/internal/usecase"
	authenticateUserUseCase "github.com/pointltd/organization/internal/usecase/auth"
	createUserUseCase "github.com/pointltd/organization/internal/usecase/user"
	"log/slog"
)

type serviceProvider struct {
	db *pgxpool.Pool

	log *slog.Logger

	userMapper mapper.UserMapper

	userRepository repository.UserRepository

	authenticateUserUseCase usecase.AuthenticateUserUseCase
	createUserUseCase       usecase.CreateUserUseCase
	listUsersUseCase        usecase.ListUsersUseCase

	authController controller.AuthController
	userController controller.UserController
}

func newServiceProvider(db *pgxpool.Pool, logger *slog.Logger) *serviceProvider {
	return &serviceProvider{
		db:  db,
		log: logger,
	}
}

func (s *serviceProvider) UserMapper() mapper.UserMapper {
	if s.userMapper == nil {
		s.userMapper = userMapper.NewUserMapper()
	}

	return s.userMapper
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.db, s.UserMapper())
	}

	return s.userRepository
}

func (s *serviceProvider) AuthenticateUserUseCase() usecase.AuthenticateUserUseCase {
	if s.authenticateUserUseCase == nil {
		s.authenticateUserUseCase = authenticateUserUseCase.NewAuthenticateUserUseCase(s.UserRepository())
	}

	return s.authenticateUserUseCase
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

func (s *serviceProvider) AuthController() controller.AuthController {
	if s.authController == nil {
		s.authController = authController.NewAuthController(s.AuthenticateUserUseCase())
	}

	return s.authController
}

func (s *serviceProvider) UserController() controller.UserController {
	if s.userController == nil {
		s.userController = userController.NewController(s.CreateUserUseCase(), s.ListUsersUseCase())
	}

	return s.userController
}
