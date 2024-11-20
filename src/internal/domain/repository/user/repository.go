package user

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pointltd/organization/internal/domain/entity"
	def "github.com/pointltd/organization/internal/domain/repository"
	"log"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]entity.User, error) {
	rows, err := r.db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []entity.User
	for rows.Next() {
		var user entity.User
		var firstName, createdById, updatedById, deletedById sql.NullString

		err = rows.Scan(
			&user.ID,
			&user.Password,
			&user.Contacts.Email,
			&user.Contacts.Phone,
			&firstName,
			&user.Info.LastName,
			&user.Timestamp.CreatedAt,
			&user.Timestamp.UpdatedAt,
			&user.Timestamp.DeletedAt,
			&createdById,
			&updatedById,
			&deletedById,
		)
		if err != nil {
			log.Fatalf(err.Error())
		}

		if firstName.Valid {
			user.Info.FirstName = firstName.String
		}
		if createdById.Valid {
			user.UserStamp.CreatedById = createdById.String
		}
		if updatedById.Valid {
			user.UserStamp.UpdatedById = updatedById.String
		}
		if deletedById.Valid {
			user.UserStamp.DeletedById = deletedById.String
		}

		users = append(users, user)
		log.Println("User: ", user)
	}
	return users, nil
}

func (r *repository) Save(user entity.User) error {
	return nil
}
