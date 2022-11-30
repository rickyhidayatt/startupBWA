package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:R1zky132@tcp(127.0.0.1:3306)/bwastart?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	campaginRepository := campaign.NewRepository(db)
	campaigns, err := campaginRepository.FindByUserID(1)

	fmt.Println("Debug")

	fmt.Println("Debug")
	fmt.Println("Debug")
	fmt.Println(len(campaigns))

	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)

		if len(campaign.CampaignImages) > 0 {

			fmt.Println(campaign.CampaignImages[0].FileName)
		}

	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.ChekEmailAvaliable)
	api.POST("/avatars", authMiddelware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}

// func ini buat auth jadi kalo kita masukin url localhost 8080/avatar gak langsung ke url  avatar, tapi di autentifikasi dulu pake func ini, kalo aman baru lanjut
func authMiddelware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIresponse("Unauthorized-1", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//Buat ngambil token nya aja
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidasiToken(tokenString)

		if err != nil {
			response := helper.APIresponse("Unauthorized 2", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims) //ngambil data yang ada di token lalu di ubah ke jwt

		if !ok || !token.Valid {
			response := helper.APIresponse("Unauthorized 3", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64)) //by default jwt kalo angka masuk ke jwt token jadi float 64
		userr, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIresponse("Unauthorized 4", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", userr) //current user maksudnya user yang sedang login dan parameternya user
	}
}
