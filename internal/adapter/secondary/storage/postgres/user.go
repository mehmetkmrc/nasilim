package postgres

import (
	"database/sql"
	"fmt"

	"github.com/mehmetkmrc/nasilim.git/internal/core/domain/model"
)

type PostgresUserRepository struct {
	DB *sql.DB
}



func NewPostgresUserRepository(connString string) (*PostgresUserRepository, error) {
	db, err := sql.Open("pgx", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()
	return &PostgresUserRepository{DB: db}, nil
}


func (r *PostgresUserRepository) Create(user *model.User) error {
	query := `
		INSERT INTO (email, password, first_name, birth_date)
		VALUES($1, $2, $3, $4)
	`
	//run the query
	_, err := r.DB.Exec(query, user.Email, user.Password, user.FirstName, user.BirthDate)
	if err != nil{
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}