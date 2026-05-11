package library

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

func (r *Repository) Add(userID, mangaID int, status string) error {
	_, err := r.db.Exec(
		`INSERT OR IGNORE INTO user_manga_library (user_id, manga_id, status)
         VALUES (?, ?, ?)`,
		userID, mangaID, status,
	)
	return err
}

func (r *Repository) GetByUser(userID int) ([]models.UserMangaLibrary, error) {
	rows, err := r.db.Query(
		`SELECT id, user_id, manga_id, status 
         FROM user_manga_library 
         WHERE user_id = ?`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.UserMangaLibrary

	for rows.Next() {
		var item models.UserMangaLibrary

		if err := rows.Scan(
			&item.ID,
			&item.UserID,
			&item.MangaID,
			&item.Status,
		); err != nil {
			return nil, err
		}

		list = append(list, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

// REMOVE FROM LIBRARY
func (r *Repository) Remove(userID, mangaID int) error {
	_, err := r.db.Exec(
		`DELETE FROM user_manga_library 
		 WHERE user_id = ? AND manga_id = ?`,
		userID, mangaID,
	)
	return err
}

// UPDATE STATUS
func (r *Repository) UpdateStatus(userID, mangaID int, status string) error {
	_, err := r.db.Exec(
		`UPDATE user_manga_library 
		 SET status = ?
		 WHERE user_id = ? AND manga_id = ?`,
		status, userID, mangaID,
	)
	return err
}
