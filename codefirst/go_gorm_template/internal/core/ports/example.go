package ports

import "gorm_template/internal/core/domain"

type ExamplePort interface {
	Create(example *domain.Example) error
	Update(example *domain.Example) error
	Delete(id uint) error
	FindByID(id uint) (*domain.Example, error)
	FindAll() ([]domain.Example, error)
}
