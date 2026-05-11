package library

import "mangahub/pkg/models"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddToLibrary(userID, mangaID int, status string) error {
	return s.repo.Add(userID, mangaID, status)
}

func (s *Service) GetLibrary(userID int) ([]models.UserMangaLibrary, error) {
	return s.repo.GetByUser(userID)
}

// REMOVE FROM LIBRARY
func (s *Service) RemoveFromLibrary(userID, mangaID int) error {
	return s.repo.Remove(userID, mangaID)
}

// UPDATE STATUS
func (s *Service) UpdateStatus(userID, mangaID int, status string) error {
	return s.repo.UpdateStatus(userID, mangaID, status)
}
