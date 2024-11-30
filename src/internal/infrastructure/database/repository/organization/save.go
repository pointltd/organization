package organization

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

func (r *repository) Save(organization entity.Organization) (entity.Organization, error) {
	id, err := uuid.NewV7()
	if err != nil {
		r.log.Error(fmt.Sprintf("error generating organization's UUID: %v", err))
		return organization, err
	}

	organization.ID = id.String()
	query := `INSERT INTO organizations (id, name, owner_id, created_at, updated_at, deleted_at) 
			VALUES (@id, @name, @owner_id, @created_at, @updated_at, @deleted_at)
			RETURNING *`

	args := r.organizationMapper.MapEntityToArg(organization)

	row, err := r.db.Query(context.Background(), query, args)

	if err != nil {
		r.log.Error(fmt.Sprintf("error saving organization: %v", err))
		return organization, err
	}

	defer row.Close()

	organizationModel, err := pgx.CollectOneRow(row, pgx.RowToStructByName[model.Organization])

	if err != nil {
		r.log.Error(fmt.Sprintf("error mapping organization model: %v", err))
		return organization, err
	}

	organization = r.organizationMapper.MapModelToEntity(organizationModel)

	return organization, nil
}
