package organization

import (
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/domain/repository"
	_def "github.com/pointltd/organization/internal/usecase"
	"log/slog"
	"time"
)

var _ _def.CreateOrganizationUseCase = (*createOrganizationUseCase)(nil)

type createOrganizationUseCase struct {
	log                    *slog.Logger
	organizationRepository repository.OrganizationRepository
}

func NewCreateOrganizationUseCase(
	log *slog.Logger,
	organizationRepository repository.OrganizationRepository,
) *createOrganizationUseCase {
	return &createOrganizationUseCase{
		log:                    log,
		organizationRepository: organizationRepository,
	}
}

func (u createOrganizationUseCase) Execute(name string, ownerId string) (entity.Organization, error) {
	organization := entity.Organization{
		Name:    name,
		OwnerID: ownerId,
	}

	timestamps := entity.Timestamp{
		CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		UpdatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
	}

	organization.Timestamp = timestamps

	organization, err := u.organizationRepository.Save(organization)
	if err != nil {
		return entity.Organization{}, err
	}

	return organization, nil
}
