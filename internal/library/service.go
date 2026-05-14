package library

import (
	"fmt"
	"mangahub/internal/manga"
	"mangahub/internal/progress"
	"mangahub/pkg/models"
)

type Service struct {
	repo         *Repository
	progressRepo *progress.Repository
	mangaRepo    *manga.Repository
}

func NewService(
	repo *Repository,
	progressRepo *progress.Repository,
	mangaRepo *manga.Repository,
) *Service {

	return &Service{
		repo:         repo,
		progressRepo: progressRepo,
		mangaRepo:    mangaRepo,
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

func (s *Service) UpdateStatus(userID, mangaID int, status string) error {
	return s.repo.UpdateStatus(userID, mangaID, status)
}

// Create progress
func (s *Service) Create(userID, mangaID int) error {

	_, err := s.mangaRepo.GetByID(mangaID)
	if err != nil {
		return err
	}

	err = s.progressRepo.Create(userID, mangaID, 1)
	if err != nil {
		return err
	}

	return nil
}

// Update progress
func (s *Service) Update(userID, mangaID, chapter int) error {

	manga, err := s.mangaRepo.GetByID(mangaID)
	if err != nil {
		return err
	}

	if chapter > manga.TotalChapters {
		return fmt.Errorf("chapter exceeds total chapters")
	}

	return s.progressRepo.Update(userID, mangaID, chapter)
}

// Get progress
func (s *Service) GetProgress(userID, mangaID int) (interface{}, error) {

	progress, err := s.progressRepo.GetProgress(userID, mangaID)
	if err != nil {
		return nil, err
	}

	return progress, nil
}
