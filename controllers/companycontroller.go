package controllers

import (
	"errors"
	"net/http"
	"strings"
	"golang.org/x/crypto/bcrypt"

	"golang-erp/models"
	"golang-erp/validators"
	"golang-erp/repositories"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
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
			c.JSON(http.StatusBadRequest, gin.H{"login": nil, "result": errFind.Error()})
			return
		}
		hashPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))
		if hashPassword != nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "Bad Login!"})
			return
		}
		signed := SignHash("234423", "#ads34!^%", user.Email)
		md5 := CalculateMD5(signed)
		c.JSON(200, gin.H{"login": user.Email, "hash": signed, "md5": md5, "result": "ok!"})
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

/**
 * Compute HMAC
 */
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

/**
 * Compute MD5
 */
func CalculateMD5(content string) string {
	h := md5.New()
	h.Write([]byte(content))
    return hex.EncodeToString(h.Sum(nil))
}

/**
 * Signing Signature for Secure Request.
 * SignHash("234423", "#ads34!^%", "your content")
 */
func SignHash(apiKey string, apiSecret string, content string) string {
	replaceApi := strings.TrimSpace(apiKey)
	replaceSecret := strings.TrimSpace(apiSecret)
	bodymd5 := CalculateMD5(content)
	stringToSign :=
		"AUTH/auth?api_key=" + strings.ToLower(replaceApi) + "&api_secret=" + strings.ToLower(replaceSecret) + "&auth_version=1.0" + "&body_md5=" + bodymd5
	return ComputeHmac256(stringToSign, apiSecret)
}

// Samples
func (s *service) FindAllCompany() []models.Company {
	var result []models.Company
	for _, c := range s.companyRepo.FindAll() {
		result = append(result, assemble(c))
	}
	return result
}
