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

// Get Library
func (r *Repository) GetByUser(userID int) ([]models.LibraryItem, error) {
	rows, err := r.db.Query(`
		SELECT
			m.id,
			m.title,
			l.status,
			p.current_chapter
		FROM user_manga_library l
		JOIN manga m
			ON l.manga_id = m.id
		JOIN user_progress p
			ON p.user_id = l.user_id
			AND p.manga_id = l.manga_id
		WHERE l.user_id = ?
    `, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.LibraryItem

	for rows.Next() {
		var item models.LibraryItem

		err := rows.Scan(
			&item.MangaID,
			&item.Title,
			&item.Status,
			&item.CurrentChapter,
		)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
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
