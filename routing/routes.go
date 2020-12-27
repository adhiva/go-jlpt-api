package routing

import (
	"go-jlpt-n5/app/healthcheck"
	"go-jlpt-n5/config"

	HCHandler "go-jlpt-n5/app/healthcheck/handler"

	"github.com/gin-gonic/gin"
)

var basePath = config.Config.App.BasePath

func HealthCheckHttpHandler(r *gin.Engine, us healthcheck.IHealthCheckUsecase) {
	handler := &HCHandler.HealthCheckHandler{
		HealthCheckUsecase: us,
	}
	route := r.Group(basePath)
	route.GET("/check", handler.Check)
}
