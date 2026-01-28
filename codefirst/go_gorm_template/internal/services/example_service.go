package services

import (
	"gorm_template/internal/core/domain"
	"gorm_template/internal/core/ports"
)

type ExampleService struct {
	repo ports.ExamplePort
}

func NewExampleService(repo ports.ExamplePort) *ExampleService {
	return &ExampleService{repo}
}

func (s *ExampleService) CreateExample(example *domain.Example) error {
	return s.repo.Create(example)
}

func (s *ExampleService) UpdateExample(example *domain.Example) error {
	return s.repo.Update(example)
}

func (s *ExampleService) DeleteExample(id uint) error {
	return s.repo.Delete(id)
}

func (s *ExampleService) GetExampleByID(id uint) (*domain.Example, error) {
	return s.repo.FindByID(id)
}

func (s *ExampleService) GetExamples() ([]domain.Example, error) {
	return s.repo.FindAll()
}
