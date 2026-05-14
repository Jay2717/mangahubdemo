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

// create manga
func (r *Repository) Create(m models.Manga) error {

	query := `
		INSERT INTO manga(
			title,
			author,
			status,
			total_chapter,
			description,
			genres,
			cover_url
		)
		VALUES(?, ?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(
		query,
		m.Title,
		m.Author,
		m.Status,
		m.Description,
		m.Genres,
		m.CoverURL,
	)

	return err
}

// Get all manga
func (r *Repository) GetAll() ([]models.Manga, error) {

	rows, err := r.db.Query(`
		SELECT 
			id,
			title,
			author,
			status,
			description,
			genres,
			cover_url
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
			&m.CoverURL,
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

// Get manga by id
func (r *Repository) GetByID(id int) (models.Manga, error) {

	var m models.Manga

	query := `
		SELECT 
			id,
			title,
			author,
			status,
			description,
			genres,
			cover_url
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
		&m.CoverURL,
	)

	return m, err
}

// Search manga
func (r *Repository) Search(
	query string,
	author string,
	status string,
	genre string,
) ([]models.Manga, error) {

	sqlQuery := `
		SELECT
			id,
			title,
			author,
			description,
			status,
			genres,
			cover_url
		FROM manga
		WHERE 1=1
	`

	var args []interface{}

	// SEARCH TITLE
	if query != "" {
		sqlQuery += `
			AND REPLACE(LOWER(title), ' ', '')
			LIKE REPLACE(LOWER(?), ' ', '')
		`
		args = append(args, "%"+query+"%")
	}

	// SEARCH AUTHOR
	if author != "" {
		sqlQuery += `
			AND REPLACE(LOWER(author), ' ', '')
			LIKE REPLACE(LOWER(?), ' ', '')
		`
		args = append(args, "%"+author+"%")
	}

	// FILTER STATUS
	if status != "" {
		sqlQuery += `
			AND LOWER(status) = LOWER(?)
		`
		args = append(args, status)
	}

	// FILTER GENRE
	if genre != "" {
		sqlQuery += `
			AND LOWER(genres) LIKE LOWER(?)
		`
		args = append(args, "%"+genre+"%")
	}

	rows, err := r.db.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mangas := make([]models.Manga, 0)

	for rows.Next() {
		var manga models.Manga

		err := rows.Scan(
			&manga.ID,
			&manga.Title,
			&manga.Author,
			&manga.Description,
			&manga.Status,
			&manga.Genres,
			&manga.CoverURL,
		)

		if err != nil {
			return nil, err
		}

		mangas = append(mangas, manga)
	}

	return mangas, nil
}
