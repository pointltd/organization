package point

import (
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/domain/repository"
	_def "github.com/pointltd/organization/internal/usecase"
	"log/slog"
)

var _ _def.CreatePointUseCase = (*createPointUseCase)(nil)

type createPointUseCase struct {
	pointRepository repository.PointRepository
	log             *slog.Logger
}

func NewCreatePointUseCase(pointRepository repository.PointRepository, log *slog.Logger) *createPointUseCase {
	return &createPointUseCase{
		pointRepository: pointRepository,
		log:             log,
	}
}

func (u createPointUseCase) Execute(name string, organizationId string) (entity.Point, error) {
	point := entity.Point{
		Name:           name,
		OrganizationId: organizationId,
	}

	point, err := u.pointRepository.Save(point)
	if err != nil {
		return entity.Point{}, err
	}

	return point, nil
}
