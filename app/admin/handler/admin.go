package handler

import (
	"fmt"
	AdminInterface "go-jlpt-n5/app/admin"
	"go-jlpt-n5/libraries"
	API "go-jlpt-n5/libraries/api"
	"go-jlpt-n5/request"

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

func (a *AdminHandler) StoreAdmin(c *gin.Context) {
	var (
		req *request.Admin
		err error
	)

	// Bind Json To Struct Request
	err = c.BindJSON(&req)
	if err != nil {
		fmt.Println("Error Log : ", err)
	}

	isValid, errorMessage := libraries.RequestValidation(req)

	if !isValid {
		API.ResponseFailValidation(c, errorMessage)
	}

	// Store Struct to DB
	admin, err := a.AdminUsecase.Store(req)
	fmt.Println("Check Admin : ", admin)

	return
}
