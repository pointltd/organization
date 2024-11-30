package organization

import (
	"github.com/jackc/pgx/v5/pgxpool"
	_def "github.com/pointltd/organization/internal/domain/repository"
	"github.com/pointltd/organization/internal/infrastructure/database/mapper"
	"log/slog"
)

var _ _def.OrganizationRepository = (*repository)(nil)

type repository struct {
	db                 *pgxpool.Pool
	organizationMapper mapper.OrganizationMapper
	log                *slog.Logger
}

func NewOrganizationRepository(db *pgxpool.Pool, organizationMapper mapper.OrganizationMapper, log *slog.Logger) *repository {
	return &repository{
		db:                 db,
		organizationMapper: organizationMapper,
		log:                log,
	}
}
