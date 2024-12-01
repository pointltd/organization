package point

import (
	"github.com/jackc/pgx/v5/pgxpool"
	def "github.com/pointltd/organization/internal/domain/repository"
	"github.com/pointltd/organization/internal/infrastructure/database/mapper"
	"log/slog"
)

var _ def.PointRepository = (*repository)(nil)

type repository struct {
	db          *pgxpool.Pool
	pointMapper mapper.PointMapper
	log         *slog.Logger
}

func NewPointRepository(db *pgxpool.Pool, pointMapper mapper.PointMapper, log *slog.Logger) *repository {
	return &repository{
		db:          db,
		pointMapper: pointMapper,
		log:         log,
	}
}
