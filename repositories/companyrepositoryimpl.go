package repositories

import (
	"golang-erp/models"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

type CompanyRepositoryImpl struct {
	Db *gorm.DB
}

func (r *CompanyRepositoryImpl) FindAll() []*models.CompanyResults {
	var result []*models.CompanyResults
	r.Db.Raw("SELECT id, name, email, telephone, code, is_active, is_verified FROM companies").Scan(&result)
	return result
}
