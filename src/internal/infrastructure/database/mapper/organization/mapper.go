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

	var userStamp = entity.UserStamp{}

	if model.CreatedById.Valid {
		userStamp.CreatedById = &model.CreatedById.String
	}
	if model.UpdatedById.Valid {
		userStamp.UpdatedById = &model.UpdatedById.String
	}
	if model.DeletedById.Valid {
		userStamp.DeletedById = &model.DeletedById.String
	}

	var organization = entity.Organization{
		ID:        model.ID,
		Name:      model.Name,
		OwnerID:   model.OwnerID,
		Timestamp: timestamp,
		UserStamp: userStamp,
	}

	return organization
}

func (m *organizationMapper) MapEntityToArg(organization entity.Organization) pgx.NamedArgs {
	args := pgx.NamedArgs{
		"id":       organization.ID,
		"name":     organization.Name,
		"owner_id": organization.OwnerID,
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

	if organization.UserStamp.CreatedById != nil {
		args["created_by_id"] = organization.UserStamp.CreatedById
	}

	if organization.UserStamp.UpdatedById != nil {
		args["updated_by_id"] = organization.UserStamp.UpdatedById
	}

	if organization.UserStamp.DeletedById != nil {
		args["deleted_by_id"] = organization.UserStamp.DeletedById
	}

	return args
}
