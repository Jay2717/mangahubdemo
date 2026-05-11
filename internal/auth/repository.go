package auth

import (
	"mangahub/pkg/database"
	"mangahub/pkg/models"
)

func CreateUser(u models.User) error {
	_, err := database.DB.Exec(
		"INSERT INTO users(id, username, password_hash) VALUES (?, ?, ?)",
		u.ID, u.Username, u.PasswordHash,
	)
	return err
}

func GetUserByUsername(username string) (models.User, error) {
	row := database.DB.QueryRow(
		"SELECT id, username, password_hash FROM users WHERE username = ?",
		username,
	)

	var u models.User
	err := row.Scan(&u.ID, &u.Username, &u.PasswordHash)
	return u, err
}