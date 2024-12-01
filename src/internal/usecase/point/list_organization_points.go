package point

import (
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/domain/repository"
	def "github.com/pointltd/organization/internal/usecase"
)

var _ def.ListOrganizationPointsUseCase = (*listOrganizationPointsUseCase)(nil)

type listOrganizationPointsUseCase struct {
	repository repository.OrganizationRepository
}

func NewListOrganizationPointsUseCase(
	repository repository.OrganizationRepository,
) *listOrganizationPointsUseCase {
	return &listOrganizationPointsUseCase{
		repository: repository,
	}
}

func (u listOrganizationPointsUseCase) Execute(organizationId string) ([]entity.Point, error) {
	return u.repository.GetPoints(organizationId)
}
