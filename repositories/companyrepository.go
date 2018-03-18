package repositories

import (
	"golang-erp/models"

	_ "github.com/lib/pq"
)

type CompanyRepository interface {
	Store(company *models.Company) error
	FindCode(email models.CompanyEmail) (*models.Company, error)
	FindAll() []*models.Company
}
