package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

func (r *repository) GetOrganizations(userId string) ([]entity.Organization, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM organizations WHERE owner_id = $1", userId)
	if err != nil {
		r.log.Error(fmt.Sprintf("failed to get user's organizations: %s", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var organizations = make([]entity.Organization, 0)

	all, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Organization])

	if err != nil {
		r.log.Error(fmt.Sprintf("failed to map organization row to entity: %s", err.Error()))
		return nil, err
	}

	for _, row := range all {
		organization := r.organizationMapper.MapModelToEntity(row)

		organizations = append(organizations, organization)
	}

	return organizations, nil
}
