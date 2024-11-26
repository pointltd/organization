package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pointltd/organization/internal/domain/entity"
	def "github.com/pointltd/organization/internal/domain/repository"
	"github.com/pointltd/organization/internal/infrastructure/database/mapper"
	"github.com/pointltd/organization/internal/infrastructure/database/model"
	"log"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	db         *pgxpool.Pool
	userMapper mapper.UserMapper
}

func NewUserRepository(db *pgxpool.Pool, userMapper mapper.UserMapper) *repository {
	return &repository{
		db:         db,
		userMapper: userMapper,
	}
}

func (r *repository) FindByEmail(email string) (*entity.User, error) {
	row, err := r.db.Query(context.Background(), "SELECT * FROM users WHERE email = $1 LIMIT 1", email)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	userModel, err := pgx.CollectOneRow(row, pgx.RowToStructByName[model.User])
	if err != nil {
		return nil, err
	}

	user := r.userMapper.MapModelToEntity(userModel)

	return &user, nil
}

func (r *repository) GetAll() ([]entity.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users = make([]entity.User, 0)

	all, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])

	if err != nil {
		log.Fatal(err)
	}

	for _, row := range all {
		user := r.userMapper.MapModelToEntity(row)

		users = append(users, user)
	}

	return users, nil
}

func (r *repository) Save(user entity.User) (entity.User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		log.Fatal(fmt.Sprintf("error generating user UUID: %v\n", err))
	}

	user.ID = id.String()

	query :=
		`INSERT INTO users (id, password, first_name, last_name, email, phone, created_at, updated_at, deleted_at, created_by_id, updated_by_id, deleted_by_id) 
		VALUES (@id, @password, @first_name, @last_name, @email, @phone, @created_at, @updated_at, @deleted_at, @created_by_id, @updated_by_id, @deleted_by_id) 
		RETURNING *`

	args := r.userMapper.MapEntityToArg(user)

	row, err := r.db.Query(context.Background(), query, args)

	if err != nil {
		return user, err
	}

	defer row.Close()

	userModel, err := pgx.CollectOneRow(row, pgx.RowToStructByName[model.User])
	if err != nil {
		return user, err
	}

	user = r.userMapper.MapModelToEntity(userModel)

	return user, nil
}
