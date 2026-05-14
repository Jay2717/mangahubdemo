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

func (s *Service) GetByID(id int) (models.Manga, error) {
	return s.repo.GetByID(id)
}

func (s *Service) SearchManga(
	query string,
	author string,
	status string,
	genre string,
) ([]models.Manga, error) {

	return s.repo.Search(
		query,
		author,
		status,
		genre,
	)
}
