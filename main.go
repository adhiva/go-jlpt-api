package main

import (
	"fmt"
	"log"
	"os"

	HCRepository "go-jlpt-n5/app/healthcheck/repository"
	HCUseCase "go-jlpt-n5/app/healthcheck/usecase"
	"go-jlpt-n5/config"
	gorm "go-jlpt-n5/db"
	routes "go-jlpt-n5/routing"

	"github.com/gin-gonic/gin"
)

var appConfig = config.Config.App

func init() {
	os.Setenv("TZ", appConfig.TimeZone)
	fmt.Println("TZ : ", appConfig.TimeZone)
	// raven.SetDSN(config.Config.SENTRY.DSN)
}

func main() {
	if appConfig.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Main program.
	r := gin.New()
	db := gorm.MysqlConn()
	hcr := HCRepository.NewHealthCheckRepository(db)
	hcu := HCUseCase.NewHealthCheckUsecase(hcr)
	routes.HealthCheckHttpHandler(r, hcu)
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// CORS Middleware
	r.Use(CORSMiddleware())
	if err := r.Run(fmt.Sprintf(":%s", appConfig.Port)); err != nil {
		log.Fatal(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-MA-Client, X-Platform, X-Api-Key, X-Secret-Key, Accept-Language, X-Product, X-Payment-Token, X-Request-Time")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
