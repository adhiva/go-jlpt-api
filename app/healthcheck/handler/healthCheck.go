package handler

import (
	HealthCheckInterface "go-jlpt-n5/app/healthcheck"
	API "go-jlpt-n5/libraries/api"
	response "go-jlpt-n5/response"
	"go-jlpt-n5/config"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
	HealthCheckUsecase HealthCheckInterface.IHealthCheckUsecase
}

func (a *HealthCheckHandler) Check(c *gin.Context) {
	healthCheck := a.HealthCheckUsecase.GetDBTimestamp()
	res := &response.HealthCheckResponse{
		HealthStatus: "GOOD ",
		DBTimestamp:  healthCheck.CurrentTimestamp,
		Environment: config.Config.App.ENV,
	}

	API.ResponseJSON(c, res)
	return
}
