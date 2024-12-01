package organization

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

func (r *repository) GetPoints(organizationId string) ([]entity.Point, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM points WHERE organization_id = $1", organizationId)
	if err != nil {
		r.log.Error(fmt.Sprintf("failed to get organization's point: %s", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var points = make([]entity.Point, 0)

	all, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.Point])

	if err != nil {
		r.log.Error(fmt.Sprintf("failed to map point row to entity: %s", err.Error()))
		return nil, err
	}

	for _, row := range all {
		point := r.pointMapper.MapModelToEntity(row)

		points = append(points, point)
	}

	return points, nil
}
