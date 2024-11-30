package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pointltd/organization/internal/domain/repository"
	"github.com/pointltd/organization/internal/infrastructure/database/mapper"
	organizationMapper "github.com/pointltd/organization/internal/infrastructure/database/mapper/organization"
	userMapper "github.com/pointltd/organization/internal/infrastructure/database/mapper/user"
	organizationRepository "github.com/pointltd/organization/internal/infrastructure/database/repository/organization"
	userRepository "github.com/pointltd/organization/internal/infrastructure/database/repository/user"
	"github.com/pointltd/organization/internal/usecase"
	authenticateUserUseCase "github.com/pointltd/organization/internal/usecase/auth"
	createOrganizationUseCase "github.com/pointltd/organization/internal/usecase/organization"
	createUserUseCase "github.com/pointltd/organization/internal/usecase/user"
	"log/slog"
)

type serviceProvider struct {
	db *pgxpool.Pool

	log *slog.Logger

	userMapper         mapper.UserMapper
	organizationMapper mapper.OrganizationMapper

	userRepository         repository.UserRepository
	organizationRepository repository.OrganizationRepository

	authenticateUserUseCase   usecase.AuthenticateUserUseCase
	createUserUseCase         usecase.CreateUserUseCase
	listUsersUseCase          usecase.ListUsersUseCase
	createOrganizationUseCase usecase.CreateOrganizationUseCase
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

func (s *serviceProvider) OrganizationMapper() mapper.OrganizationMapper {
	if s.organizationMapper == nil {
		s.organizationMapper = organizationMapper.NewOrganizationMapper()
	}

	return s.organizationMapper
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.db, s.UserMapper(), s.log)
	}

	return s.userRepository
}

func (s *serviceProvider) OrganizationRepository() repository.OrganizationRepository {
	if s.organizationRepository == nil {
		s.organizationRepository = organizationRepository.NewOrganizationRepository(s.db, s.OrganizationMapper(), s.log)
	}

	return s.organizationRepository

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

func (s *serviceProvider) CreateOrganizationUseCase() usecase.CreateOrganizationUseCase {
	if s.createOrganizationUseCase == nil {
		s.createOrganizationUseCase = createOrganizationUseCase.NewCreateOrganizationUseCase(
			s.log,
			s.OrganizationRepository(),
		)
	}

	return s.createOrganizationUseCase
}
