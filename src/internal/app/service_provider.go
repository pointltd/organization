package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pointltd/organization/internal/config"
	"github.com/pointltd/organization/internal/domain/repository"
	"github.com/pointltd/organization/internal/infrastructure/database/mapper"
	organizationMapper "github.com/pointltd/organization/internal/infrastructure/database/mapper/organization"
	pointMapper "github.com/pointltd/organization/internal/infrastructure/database/mapper/point"
	userMapper "github.com/pointltd/organization/internal/infrastructure/database/mapper/user"
	organizationRepository "github.com/pointltd/organization/internal/infrastructure/database/repository/organization"
	pointRepository "github.com/pointltd/organization/internal/infrastructure/database/repository/point"
	userRepository "github.com/pointltd/organization/internal/infrastructure/database/repository/user"
	"github.com/pointltd/organization/internal/usecase"
	authenticateUserUseCase "github.com/pointltd/organization/internal/usecase/auth"
	createOrganizationUseCase "github.com/pointltd/organization/internal/usecase/organization"
	createPointUseCase "github.com/pointltd/organization/internal/usecase/point"
	listOrganizationPointsUseCase "github.com/pointltd/organization/internal/usecase/point"
	createUserUseCase "github.com/pointltd/organization/internal/usecase/user"
	"log/slog"
)

type serviceProvider struct {
	db *pgxpool.Pool

	log       *slog.Logger
	appConfig config.AppConfig

	userMapper         mapper.UserMapper
	organizationMapper mapper.OrganizationMapper
	pointMapper        mapper.PointMapper

	userRepository         repository.UserRepository
	organizationRepository repository.OrganizationRepository
	pointRepository        repository.PointRepository

	authenticateUserUseCase       usecase.AuthenticateUserUseCase
	createUserUseCase             usecase.CreateUserUseCase
	listUsersUseCase              usecase.ListUsersUseCase
	createOrganizationUseCase     usecase.CreateOrganizationUseCase
	listUserOrganizationsUseCase  usecase.ListUserOrganizationsUseCase
	createPointUseCase            usecase.CreatePointUseCase
	listOrganizationPointsUseCase usecase.ListOrganizationPointsUseCase
}

func newServiceProvider(db *pgxpool.Pool, logger *slog.Logger, appConfig config.AppConfig) *serviceProvider {
	return &serviceProvider{
		db:        db,
		log:       logger,
		appConfig: appConfig,
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

func (s *serviceProvider) PointMapper() mapper.PointMapper {
	if s.pointMapper == nil {
		s.pointMapper = pointMapper.NewPointMapper()
	}

	return s.pointMapper
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.db, s.UserMapper(), s.OrganizationMapper(), s.log)
	}

	return s.userRepository
}

func (s *serviceProvider) OrganizationRepository() repository.OrganizationRepository {
	if s.organizationRepository == nil {
		s.organizationRepository = organizationRepository.NewOrganizationRepository(
			s.db,
			s.OrganizationMapper(),
			s.PointMapper(),
			s.log,
		)
	}

	return s.organizationRepository

}

func (s *serviceProvider) PointRepository() repository.PointRepository {
	if s.pointRepository == nil {
		s.pointRepository = pointRepository.NewPointRepository(s.db, s.PointMapper(), s.log)
	}

	return s.pointRepository
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

func (s *serviceProvider) ListUserOrganizationsUseCase() usecase.ListUserOrganizationsUseCase {
	if s.listUserOrganizationsUseCase == nil {
		s.listUserOrganizationsUseCase = createUserUseCase.NewListUserOrganizationsUseCase(
			s.UserRepository(),
			s.log,
		)
	}

	return s.listUserOrganizationsUseCase
}

func (s *serviceProvider) ListOrganizationPointsUseCase() usecase.ListOrganizationPointsUseCase {
	if s.listOrganizationPointsUseCase == nil {
		s.listOrganizationPointsUseCase = listOrganizationPointsUseCase.NewListOrganizationPointsUseCase(
			s.OrganizationRepository(),
		)
	}

	return s.listOrganizationPointsUseCase
}

func (s *serviceProvider) CreatePointUseCase() usecase.CreatePointUseCase {
	if s.createPointUseCase == nil {
		s.createPointUseCase = createPointUseCase.NewCreatePointUseCase(
			s.PointRepository(),
			s.log,
		)
	}

	return s.createPointUseCase
}
