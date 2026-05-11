package manga

import "mangahub/pkg/models"

func GetMangaList() ([]models.Manga, error) {
	return GetAllManga()
}

func CreateNewManga(m models.Manga) error {
	return CreateManga(m)
}