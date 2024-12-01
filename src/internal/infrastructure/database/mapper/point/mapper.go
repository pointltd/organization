package point

import (
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	_def "github.com/pointltd/organization/internal/infrastructure/database/mapper"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

var _ _def.PointMapper = (*pointMapper)(nil)

type pointMapper struct {
}

func NewPointMapper() *pointMapper {
	return &pointMapper{}
}

func (m *pointMapper) MapModelToEntity(model model.Point) entity.Point {
	var timestamp = entity.Timestamp{}

	if model.CreatedAt.Valid {
		timestamp.CreatedAt = &model.CreatedAt.Time
	}

	if model.UpdatedAt.Valid {
		timestamp.UpdatedAt = &model.UpdatedAt.Time
	}

	if model.DeletedAt.Valid {
		timestamp.DeletedAt = &model.DeletedAt.Time
	}

	var point = entity.Point{
		Id:             model.Id,
		Name:           model.Name,
		OrganizationId: model.OrganizationId,
		Timestamp:      timestamp,
	}

	return point
}

func (m *pointMapper) MapEntityToArg(point entity.Point) pgx.NamedArgs {
	args := pgx.NamedArgs{
		"id":              point.Id,
		"name":            point.Name,
		"organization_id": point.OrganizationId,
	}

	if point.Timestamp.CreatedAt != nil {
		args["created_at"] = point.Timestamp.CreatedAt
	}

	if point.Timestamp.UpdatedAt != nil {
		args["updated_at"] = point.Timestamp.UpdatedAt
	}

	if point.Timestamp.DeletedAt != nil {
		args["deleted_at"] = point.Timestamp.DeletedAt
	}

	return args
}
