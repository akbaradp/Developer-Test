package repository

import (
	"user_management/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Nationality{},
		&models.Customer{},
		&models.FamilyList{},
	)
	if err != nil {
		return err
	}
	return nil
}

type NationalityRepository interface {
	FindAll() ([]models.Nationality, error)
}

type nationalityRepository struct {
	db *gorm.DB
}

func NewNationalityRepository(db *gorm.DB) NationalityRepository {
	return &nationalityRepository{db}
}

func (r *nationalityRepository) FindAll() ([]models.Nationality, error) {
	var nationalities []models.Nationality
	err := r.db.Find(&nationalities).Error
	return nationalities, err
}

type CustomersRepository interface {
	FindAll() ([]models.Customer, error)
}

type customersRepository struct {
	db *gorm.DB
}

func NewCustomersRepository(db *gorm.DB) CustomersRepository {
	return &customersRepository{db}
}

func (r *customersRepository) FindAll() ([]models.Customer, error) {
	var customers []models.Customer
	err := r.db.Find(&customers).Error
	return customers, err
}

type FamilyRepository interface {
	FindAll() ([]models.FamilyList, error)
}

type familyRepository struct {
	db *gorm.DB
}

func NewFamilyRepository(db *gorm.DB) FamilyRepository {
	return &familyRepository{db}
}

func (r *familyRepository) FindAll() ([]models.FamilyList, error) {
	var families []models.FamilyList
	err := r.db.Find(&families).Error
	return families, err
}
