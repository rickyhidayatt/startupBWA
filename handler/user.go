package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		pesanError := helper.FormatErrorValidator(err)
		errorMassage := gin.H{"errors": pesanError}

		response := helper.APIresponse("Register Account Failed", http.StatusUnprocessableEntity, "Error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	c.JSON(http.StatusOK, nil)

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIresponse("Register Account Failed", http.StatusOK, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formaterr := user.FormatUser(newUser, "Tokennnntoken token disini")

	response := helper.APIresponse("Acount has been register", http.StatusOK, "Succes", formaterr)

	c.JSON(http.StatusOK, response)
}
