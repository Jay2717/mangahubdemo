package manga

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

func (r *Repository) Create(m models.Manga) error {

	query := `
		INSERT INTO manga(title, author, status, description, genres)
		VALUES(?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		m.Title,
		m.Author,
		m.Status,
		m.Description,
		m.Genres,
	)

	return err
}

func (r *Repository) GetAll() ([]models.Manga, error) {

	rows, err := r.db.Query(`
		SELECT id, title, author, status, description, genres 
		FROM manga
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.Manga

	for rows.Next() {
		var m models.Manga

		if err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.Author,
			&m.Status,
			&m.Description,
			&m.Genres,
		); err != nil {
			return nil, err
		}

		list = append(list, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *Repository) GetByID(id string) (models.Manga, error) {

	var m models.Manga

	query := `
		SELECT id, title, author, status, description, genres
		FROM manga
		WHERE id = ?
	`

	err := r.db.QueryRow(query, id).Scan(
		&m.ID,
		&m.Title,
		&m.Author,
		&m.Status,
		&m.Description,
		&m.Genres,
	)

	return m, err
}
