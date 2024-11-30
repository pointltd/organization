package user

import (
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/domain/repository"
	def "github.com/pointltd/organization/internal/usecase"
	"log/slog"
)

var _ def.ListUserOrganizationsUseCase = (*listUserOrganizationsUseCase)(nil)

type listUserOrganizationsUseCase struct {
	userRepository repository.UserRepository
	log            *slog.Logger
}

func NewListUserOrganizationsUseCase(
	userRepository repository.UserRepository,
	log *slog.Logger,
) *listUserOrganizationsUseCase {
	return &listUserOrganizationsUseCase{
		userRepository: userRepository,
		log:            log,
	}
}

func (u listUserOrganizationsUseCase) Execute(userId string) ([]entity.Organization, error) {
	return u.userRepository.GetOrganizations(userId)
}
