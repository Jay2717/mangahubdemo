package progress

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(
	userID,
	mangaID,
	chapter int,
) error {

	_, err := r.db.Exec(`
		INSERT INTO user_progress
		(user_id, manga_id, current_chapter, last_page)
		VALUES (?, ?, ?, ?)
	`, userID, mangaID, chapter, 1)

	return err
}

func (r *Repository) Delete(userID, mangaID int) error {
	_, err := r.db.Exec(`
		DELETE FROM user_progress
		WHERE user_id = ? AND manga_id = ?
	`, userID, mangaID)

	return err
}
