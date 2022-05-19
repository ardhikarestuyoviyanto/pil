package service

import (
	"pil/config"
	"pil/domain"
	"pil/internal/model"
)

type service struct {
	c    config.Config
	repo domain.AllServices
}

// CreateActivity implements domain.AllServices
func (s *service) CreateActivity(activity model.Activity) {
	s.repo.CreateActivity(activity)
}

// DeleteActivity implements domain.AllServices
func (s *service) DeleteActivity(id int) {
	s.repo.DeleteActivity(id)
}

// GetAllActivity implements domain.AllServices
func (s *service) GetAllActivity() []model.APIResponseActivity {
	return s.repo.GetAllActivity()
}

// GetByIdActivity implements domain.AllServices
func (s *service) GetByIdActivity(id int) model.APIResponseActivity {
	return s.repo.GetByIdActivity(id)
}

// UpdateActivity implements domain.AllServices
func (s *service) UpdateActivity(id int, activity model.Activity) {
	s.repo.UpdateActivity(id, activity)
}

func NewService(repo domain.AllRepository, c config.Config) domain.AllServices {
	return &service{
		c:    c,
		repo: repo,
	}
}
