package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"golang-erp/controllers"
	log "github.com/inconshreveable/log15"
	"golang-erp/models"
)

const (
	defaultPort = "8080"
)

func main() {
	var (
		addr = envString("PORT", defaultPort)
	)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	//Put Controller Here
	company := new(controllers.CompanyController)
	company.Db = db
	
	//Put All Routes In Here
	router.GET("/v2/company", company.Get)
	router.POST("/v2/company/create", company.CreateCompany)
	router.POST("/v2/company/login", company.LoginCompany)

	router.Run(":" + addr)
	srvlog := log.New("module", "app/server")
	srvlog.Info("connection open")
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
