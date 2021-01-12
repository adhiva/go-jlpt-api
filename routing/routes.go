package routing

import (
	"go-jlpt-n5/app/admin"
	"go-jlpt-n5/app/healthcheck"
	"go-jlpt-n5/config"

	AdminHandler "go-jlpt-n5/app/admin/handler"
	HCHandler "go-jlpt-n5/app/healthcheck/handler"

	"github.com/gin-gonic/gin"
)

var basePath = config.Config.App.BasePath

func HealthCheckHttpHandler(r *gin.Engine, us healthcheck.IHealthCheckUsecase) {
	handler := &HCHandler.HealthCheckHandler{
		HealthCheckUsecase: us,
	}
	route := r.Group(basePath)
	route.GET("/health-check", handler.Check)
}

func AuthHttpHandler(r *gin.Engine, us healthcheck.IHealthCheckUsecase) {
	handler := &HCHandler.HealthCheckHandler{
		HealthCheckUsecase: us,
	}
	route := r.Group(basePath)
	route.POST("/auth/login", handler.Check)
}

func AdminHttpHandler(r *gin.Engine, auc admin.IAdminUsecase) {
	handler := &AdminHandler.AdminHandler{
		AdminUsecase: auc,
	}
	route := r.Group(basePath)
	route.GET("/admins", handler.GetAdmin)
}

func PermissionHttpHandler(r *gin.Engine, us healthcheck.IHealthCheckUsecase) {
	handler := &HCHandler.HealthCheckHandler{
		HealthCheckUsecase: us,
	}
	route := r.Group(basePath)
	route.GET("/permissions", handler.Check)
	route.GET("/permission/:id", handler.Check)
	route.POST("/permission", handler.Check)
	route.DELETE("/permission/:id", handler.Check)
	route.PUT("/permission/:id", handler.Check)
}

func AdminPermissionsHttpHandler(r *gin.Engine, us healthcheck.IHealthCheckUsecase) {
	handler := &HCHandler.HealthCheckHandler{
		HealthCheckUsecase: us,
	}
	route := r.Group(basePath)
	route.GET("/admin-permissions", handler.Check)
	route.GET("/admin-permissions/:userId", handler.Check)
	route.POST("/admin-permissions/:userId", handler.Check)
	route.DELETE("/admins-permissions/:userId", handler.Check)
	route.PUT("/admins-permissions/:userId", handler.Check)
}
