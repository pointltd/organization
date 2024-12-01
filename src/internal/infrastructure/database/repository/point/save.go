package point

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pointltd/organization/internal/domain/entity"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
)

func (r *repository) Save(point entity.Point) (entity.Point, error) {
	id, err := uuid.NewV7()
	if err != nil {
		r.log.Error(fmt.Sprintf("error generating point's UUID: %v", err))
		return point, err
	}

	point.Id = id.String()

	query := `INSERT INTO points (id, name, organization_id, created_at, updated_at) 
			VALUES (@id, @name, @organization_id, @created_at, @updated_at) RETURNING *`

	args := r.pointMapper.MapEntityToArg(point)

	row, err := r.db.Query(context.Background(), query, args)
	if err != nil {
		r.log.Error(fmt.Sprintf("error saving point: %v", err))
		return point, err
	}
	defer row.Close()

	pointModel, err := pgx.CollectOneRow(row, pgx.RowToStructByName[model.Point])
	if err != nil {
		r.log.Error(fmt.Sprintf("error mapping point model: %v", err))
		return point, err
	}

	point = r.pointMapper.MapModelToEntity(pointModel)

	return point, nil
}
