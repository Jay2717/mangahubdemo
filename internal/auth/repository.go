package auth

import (
	"database/sql"
	"mangahub/pkg/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user models.User) error {

	query := `
	INSERT INTO users(username, email, password_hash)
	VALUES(?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
	)

	return err
}

func (r *Repository) GetUserByUsername(username string) (models.User, error) {
	var user models.User

	query := `
	SELECT id, username, email, password_hash FROM users
	WHERE username = ?
	`

	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	return user, err
}
