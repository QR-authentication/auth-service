package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Импорт драйвера PostgreSQL

	"github.com/QR-authentication/auth-service/internal/config"
	"github.com/QR-authentication/auth-service/internal/model"
)

type Repository struct {
	connection *sqlx.DB
}

func New(cfg *config.Config) *Repository {
	conStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Database, cfg.Postgres.Host, cfg.Postgres.Port)

	conn, err := sqlx.Connect("postgres", conStr)
	if err != nil {
		log.Fatal("error connect: ", err)
	}

	return &Repository{
		connection: conn,
	}
}

func (r *Repository) Close() {
	_ = r.connection.Close()
}

func (r *Repository) UserExists(cardNumber string) (bool, error) {
	query := `
        SELECT EXISTS (
            SELECT 1 
            FROM users 
            WHERE card_number = $1
        )`

	var exists bool
	err := r.connection.Get(&exists, query, cardNumber)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *Repository) GetUserData(cardNumber string) (*model.User, error) {
	var user model.User

	query := `
		SELECT id, card_number, name, surname, password
		FROM users
		WHERE card_number = $1`

	err := r.connection.Get(&user, query, cardNumber)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
