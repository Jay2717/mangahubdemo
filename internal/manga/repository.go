package manga

import (
	"mangahub/pkg/database"
	"mangahub/pkg/models"
)

func GetAllManga() ([]models.Manga, error) {
	rows, err := database.DB.Query("SELECT id, title, author FROM manga")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	mangas := []models.Manga{}

	for rows.Next() {
		var m models.Manga
		err := rows.Scan(&m.ID, &m.Title, &m.Author)
		if err != nil {
			return nil, err
		}
		mangas = append(mangas, m)
	}

	return mangas, nil
}

func CreateManga(m models.Manga) error {
	_, err := database.DB.Exec(
		"INSERT INTO manga(id, title, author) VALUES (?, ?, ?)",
		m.ID, m.Title, m.Author,
	)
	return err
}