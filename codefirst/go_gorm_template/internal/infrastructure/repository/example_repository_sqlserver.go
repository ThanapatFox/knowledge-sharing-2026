package repository

import (
	"gorm_template/internal/core/domain"

	"gorm.io/gorm"
)

type ExampleRepositorySQLServer struct {
	db *gorm.DB
}

func NewExampleRepositorySQLServer(db *gorm.DB) *ExampleRepositorySQLServer {
	return &ExampleRepositorySQLServer{db}
}

func (r *ExampleRepositorySQLServer) Create(example *domain.Example) error {
	return r.db.Create(example).Error
}

func (r *ExampleRepositorySQLServer) Update(example *domain.Example) error {
	return r.db.Save(example).Error
}

func (r *ExampleRepositorySQLServer) Delete(id uint) error {
	return r.db.Delete(&domain.Example{}, id).Error
}

func (r *ExampleRepositorySQLServer) FindByID(id uint) (*domain.Example, error) {
	var example domain.Example
	err := r.db.First(&example, id).Error
	return &example, err
}

func (r *ExampleRepositorySQLServer) FindAll() ([]domain.Example, error) {
	var examples []domain.Example
	err := r.db.Find(&examples).Error
	return examples, err
}
