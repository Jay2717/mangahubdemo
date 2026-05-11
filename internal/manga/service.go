package manga

import "mangahub/pkg/models"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateManga(m models.Manga) error {
	return s.repo.Create(m)
}

func (s *Service) GetAll() ([]models.Manga, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(id string) (models.Manga, error) {
	return s.repo.GetByID(id)
}
