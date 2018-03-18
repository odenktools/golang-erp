package models

import (
	"time"
	
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type CompanyEmail string

type Company struct {
	Id             uint64 `sql:"type:bigint PRIMARY KEY`
	Name           string `gorm:"type:varchar(100);not null"`
	Email          string `gorm:"type:varchar(100);not null"`
	Password       string `gorm:"type:varchar(100);not null"`
	Telephone      string `gorm:"type:varchar(15);not null"`
	Code           string `gorm:"type:varchar(15);not null"`
	Is_active      int
	Is_verified    int
	Last_login     *time.Time
	Remember_login string `gorm:"type:varchar(100);null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

type CompanyResults struct {
	Id             uint64
	Name           string
	Email          string
	Telephone      string
	Code           string
	Is_active      int
	Is_verified    int
}

type EmailNotExistsError struct{}

func (*EmailNotExistsError) Error() string {
  return "email not exists"
}

func FindCompanyByEmail(db *gorm.DB, email string) (*Company, error) {
  var company Company
  res := db.Find(&company, &Company{Email: email})
  if res.RecordNotFound() {
    return nil, &EmailNotExistsError{}
  }
  return &company, nil
}