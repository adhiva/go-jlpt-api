package handler

import (
	"fmt"
	AdminInterface "go-jlpt-n5/app/admin"
	API "go-jlpt-n5/libraries/api"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUsecase AdminInterface.IAdminUsecase
}

func (a *AdminHandler) GetAdmin(c *gin.Context) {
	admins, err := a.AdminUsecase.Get()
	if err != nil {
		fmt.Println("Error Log : ", err)
	}

	API.ResponseJSON(c, admins)
	return
}
