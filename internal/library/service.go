package library

import (
	"mangahub/internal/progress"
	"mangahub/pkg/models"
)

type Service struct {
	repo         *Repository
	progressRepo *progress.Repository
}

func NewService(
	repo *Repository,
	progressRepo *progress.Repository,
) *Service {

	return &Service{
		repo:         repo,
		progressRepo: progressRepo,
	}
}

func (s *Service) AddToLibrary(
	userID,
	mangaID int,
	status string,
	currentChapter int,
) error {

	err := s.repo.Add(userID, mangaID, status)
	if err != nil {
		return err
	}

	err = s.progressRepo.Create(
		userID,
		mangaID,
		currentChapter,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetLibrary(userID int) ([]models.LibraryItem, error) {
	return s.repo.GetByUser(userID)
}

// REMOVE FROM LIBRARY
func (s *Service) RemoveFromLibrary(userID, mangaID int) error {

	err := s.repo.Remove(userID, mangaID)
	if err != nil {
		return err
	}

	err = s.progressRepo.Delete(userID, mangaID)
	if err != nil {
		return err
	}

	return nil
}

// UPDATE STATUS
func (s *Service) UpdateStatus(userID, mangaID int, status string) error {
	return s.repo.UpdateStatus(userID, mangaID, status)
}
