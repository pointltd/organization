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

func NewRepository(db *pgxpool.Pool, userMapper mapper.UserMapper) *repository {
	return &repository{
		db:         db,
		userMapper: userMapper,
	}
}

func (r *repository) GetAll() ([]entity.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []entity.User

	all, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])

	if err != nil {
		log.Fatal(err)
	}

	for _, row := range all {
		user, err := r.userMapper.MapModelToEntity(row)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *repository) Save(user entity.User) error {
	id, err := uuid.NewV7()
	if err != nil {
		log.Fatal(fmt.Sprintf("error generating user UUID: %v\n", err))
	}

	user.ID = id.String()

	query := `INSERT INTO users (id, password, first_name, last_name, email) VALUES ($1, $2, $3, $4, $5)`

	_, err = r.db.Exec(context.Background(), query, id, user.Password, user.Info.FirstName, user.Info.LastName, user.Contacts.Email)

	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}

	return nil
}
