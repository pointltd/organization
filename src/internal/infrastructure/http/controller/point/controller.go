package point

import (
	_def "github.com/pointltd/organization/internal/infrastructure/http/controller"
	"github.com/pointltd/organization/internal/usecase"
	"log/slog"
)

var _ _def.PointController = (*controller)(nil)

type controller struct {
	createPointUseCase usecase.CreatePointUseCase
	log                *slog.Logger
}

func NewPointController(
	createPointUseCase usecase.CreatePointUseCase,
	log *slog.Logger,
) *controller {
	return &controller{
		log:                log,
		createPointUseCase: createPointUseCase,
	}
}
