package organization

import (
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	_def "github.com/pointltd/organization/internal/infrastructure/database/mapper"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

var _ _def.OrganizationMapper = (*organizationMapper)(nil)

type organizationMapper struct{}

func NewOrganizationMapper() *organizationMapper {
	return &organizationMapper{}
}

func (m *organizationMapper) MapModelToEntity(model model.Organization) entity.Organization {
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

	var organization = entity.Organization{
		Id:        model.Id,
		Name:      model.Name,
		OwnerId:   model.OwnerId,
		Timestamp: timestamp,
	}

	return organization
}

func (m *organizationMapper) MapEntityToArg(organization entity.Organization) pgx.NamedArgs {
	args := pgx.NamedArgs{
		"id":       organization.Id,
		"name":     organization.Name,
		"owner_id": organization.OwnerId,
	}

	if organization.Timestamp.CreatedAt != nil {
		args["created_at"] = organization.Timestamp.CreatedAt
	}

	if organization.Timestamp.UpdatedAt != nil {
		args["updated_at"] = organization.Timestamp.UpdatedAt
	}

	if organization.Timestamp.DeletedAt != nil {
		args["deleted_at"] = organization.Timestamp.DeletedAt
	}

	return args
}
