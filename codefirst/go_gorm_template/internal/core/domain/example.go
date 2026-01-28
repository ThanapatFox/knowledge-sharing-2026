package domain

type Example struct {
	ID   uint   `gorm:"primaryKey"`
	Code string `gorm:"unique;not null"`
	Name string `gorm:"not null"`
}
