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

func (r *Repository) Update(userID, mangaID, chapter int) error {

	res, err := r.db.Exec(`
		UPDATE user_progress
		SET current_chapter = ?, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = ? AND manga_id = ?
	`, chapter, userID, mangaID)

	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// get
func (r *Repository) GetProgress(userID, mangaID int) (map[string]interface{}, error) {

	row := r.db.QueryRow(`
		SELECT user_id, manga_id, current_chapter, last_page, updated_at
		FROM user_progress
		WHERE user_id = ? AND manga_id = ?
	`, userID, mangaID)

	var uid, mid, chapter, page int
	var updatedAt string

	err := row.Scan(&uid, &mid, &chapter, &page, &updatedAt)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user_id":         uid,
		"manga_id":        mid,
		"current_chapter": chapter,
		"last_page":       page,
		"updated_at":      updatedAt,
	}, nil
}
