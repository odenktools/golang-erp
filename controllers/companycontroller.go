package controllers

import (
	"errors"
	"net/http"
	"golang.org/x/crypto/bcrypt"

	"golang-erp/models"
	"golang-erp/validators"
	"golang-erp/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

var ErrInvalidArgument = errors.New("invalid argument")

type CompanyController struct {
	Db *gorm.DB
}

type CompanyService interface {
	FindAllCompany() []models.Company
}

type service struct {
	companyRepo repositories.CompanyRepository
}

//Get All Company
func (main *CompanyController) Get(ctx *gin.Context) {
	var result []models.CompanyResults
	main.Db.Raw("SELECT id, name, email, telephone, code, is_active, is_verified FROM companies").Scan(&result)
	ctx.JSON(200, gin.H{"result": result})
}

//Create New Company
func (main *CompanyController) CreateCompany(c *gin.Context) {
	var form validators.CompanyCreateIsValid
	if err := c.ShouldBindJSON(&form); err == nil {
		passwordHash, errPass := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
		if errPass != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		user := &models.Company{Name: form.Name, Email: form.Email, Password: string(passwordHash), Telephone: form.Telephone, Code: form.Code, Is_active: 0, Is_verified: 0}
		main.Db.Create(&user)
		c.JSON(200, gin.H{"created": user, "result": "ok!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"created": nil, "result": err.Error()})
	}
}

//Login Company
func (main *CompanyController) LoginCompany(c *gin.Context) {
	var form validators.CompanyLoginIsValid
	if err := c.ShouldBindJSON(&form); err == nil {
		user, errFind := models.FindCompanyByEmail(main.Db, form.Email)
		if errFind != nil {
			c.JSON(200, gin.H{"login": nil, "result": errFind.Error()})
			return
		}
		hashPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
		if hashPassword != nil {
			c.JSON(200, gin.H{"result": "Bad Login!"})
			return
		}
		c.JSON(200, gin.H{"login": user.Email, "result": "ok!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"login": nil, "result": err.Error()})
	}
}

// Samples
func assemble(c *models.Company) models.Company {
	return models.Company{
		Name:  string(c.Name),
		Email: c.Email,
	}
}

// Samples
func (s *service) FindAllCompany() []models.Company {
	var result []models.Company
	for _, c := range s.companyRepo.FindAll() {
		result = append(result, assemble(c))
	}
	return result
}
