package organization

import (
	_def "github.com/pointltd/organization/internal/infrastructure/http/controller"
	"github.com/pointltd/organization/internal/usecase"
	"log/slog"
)

var _ _def.OrganizationController = (*controller)(nil)

type controller struct {
	createOrganizationUseCase usecase.CreateOrganizationUseCase
	log                       *slog.Logger
}

func NewOrganizationController(
	createOrganizationUseCase usecase.CreateOrganizationUseCase,
	log *slog.Logger,
) *controller {
	return &controller{
		log:                       log,
		createOrganizationUseCase: createOrganizationUseCase,
	}
}
