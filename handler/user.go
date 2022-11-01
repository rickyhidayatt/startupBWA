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

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		pesanError := helper.FormatErrorValidator(err)
		errorMassage := gin.H{"errors": pesanError}

		response := helper.APIresponse("Login Account Failed", http.StatusUnprocessableEntity, "Error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	LoginUser, err := h.userService.Login(input)

	if err != nil {
		errorMassage := gin.H{"errors": err.Error()}

		response := helper.APIresponse("Login Account Failed", http.StatusUnprocessableEntity, "Error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formaterr := user.FormatUser(LoginUser, "Tokennnntokentoken token token")

	response := helper.APIresponse("Akun Berhasil Login", http.StatusOK, "succes", formaterr)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) ChekEmailAvaliable(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		pesanError := helper.FormatErrorValidator(err)
		errorMassage := gin.H{"errors": pesanError}

		response := helper.APIresponse("Email Failed", http.StatusUnprocessableEntity, "Error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvaliable, err := h.userService.EmailAvaliable(input)
	if err != nil {
		errorMassage := gin.H{"errors": "Server Error "}
		response := helper.APIresponse("Email Failed", http.StatusUnprocessableEntity, "Error", errorMassage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"Email Avaliable": isEmailAvaliable,
	}

	var metaMassage string

	if isEmailAvaliable {
		metaMassage = "Email Avaliable"
	} else {
		metaMassage = "Email has been registered"
	}

	response := helper.APIresponse(metaMassage, http.StatusOK, "Succes", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// Input data dari user
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_Uploaded": false}
		response := helper.APIresponse("Failed to Upload Avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
	}

	// Manggil file images dari folder Images
	// + ngambil images ditambah file.Filename
	path := "images/" + file.Filename

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_Uploaded": false}
		response := helper.APIresponse("Failed to Upload Avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
	}

	userID := 1 // Dapet dari JWT Percobaan inimah hehe :)

	_, err = h.userService.SaveAvatar(userID, path)

	if err != nil {
		data := gin.H{"is_Uploaded": false}
		response := helper.APIresponse("Failed to Upload Avatar image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
	}

	// Kalo gak error ya ke upload hehe :)
	data := gin.H{"is_Uploaded": true}
	response := helper.APIresponse("Avatar Succes Upload", http.StatusOK, "succes", data)

	c.JSON(http.StatusOK, response)
}
