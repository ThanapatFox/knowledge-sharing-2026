package repository

import (
	"gorm_template/internal/core/domain"
	"log"

	"gorm.io/gorm"
)

type ExampleRepositoryPostgres struct {
	db *gorm.DB
}

func NewExampleRepositoryPostgres(db *gorm.DB) *ExampleRepositoryPostgres {
	return &ExampleRepositoryPostgres{db}
}

// use another Create for example of using transaction
// func (r *ExampleRepositoryPostgres) Create(example *domain.Example) error {
// 	return r.db.Create(example).Error
// }

// create by using transaction and defer func if tx failed
func (r *ExampleRepositoryPostgres) Create(example *domain.Example) error {
	// Start a transaction
	tx := r.db.Begin()

	// use defer to rollback in case of panic or error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Transaction rolled back due to panic:", r)
		}
	}()

	// your operations go here
	if err := tx.Create(&example).Error; err != nil {
		tx.Rollback() // Rollback here
		log.Println("Transaction rolled back due to error:", err)
		return err
	}

	// if all goes well, commit the transaction
	if err := tx.Commit().Error; err != nil {
		log.Println("Failed to commit transaction:", err)
		return err
	} else {
		log.Println("Transaction committed successfully")
		return tx.Error
	}
}

func (r *ExampleRepositoryPostgres) Update(example *domain.Example) error {
	return r.db.Save(example).Error
}

func (r *ExampleRepositoryPostgres) Delete(id uint) error {
	return r.db.Delete(&domain.Example{}, id).Error
}

func (r *ExampleRepositoryPostgres) FindByID(id uint) (*domain.Example, error) {
	var example domain.Example
	err := r.db.First(&example, id).Error
	return &example, err
}

func (r *ExampleRepositoryPostgres) FindAll() ([]domain.Example, error) {
	var examples []domain.Example
	err := r.db.Find(&examples).Error
	return examples, err
}
