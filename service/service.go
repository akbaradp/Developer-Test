package service

import (
	"user_management/models"
	"user_management/repository"
)

type NationalityService interface {
	GetAllNationalities() ([]models.Nationality, error)
}

type nationalityService struct {
	repo repository.NationalityRepository
}

func NewNationalityService(repo repository.NationalityRepository) NationalityService {
	return &nationalityService{repo}
}

func (s *nationalityService) GetAllNationalities() ([]models.Nationality, error) {
	return s.repo.FindAll()
}

type CustomersService interface {
	GetAllCustomers() ([]models.Customer, error)
}

type customersService struct {
	repo repository.CustomersRepository
}

func NewCustomersService(repo repository.CustomersRepository) CustomersService {
	return &customersService{repo}
}

func (s *customersService) GetAllCustomers() ([]models.Customer, error) {
	return s.repo.FindAll()
}

type FamilyService interface {
	GetAllFamilies() ([]models.FamilyList, error)
}

type familyService struct {
	repo repository.FamilyRepository
}

func NewFamilyService(repo repository.FamilyRepository) FamilyService {
	return &familyService{repo}
}

func (s *familyService) GetAllFamilies() ([]models.FamilyList, error) {
	return s.repo.FindAll()
}
