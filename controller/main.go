package main

import (
	"admins/auth"
	"admins/docs"
	"admins/service"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// SaveAdmin godoc
// @Summary Endpoint used to save an admin
// @Schemes
// @Description saves an admin on the database
// @Tags admin
// @Accept json
// @Produce json
// @Param email query string true "Email of the admin" Format(email)
// @Param password query string true "Password of the admin" Format(password)
// @Success 200 {string} Admin saved
// @Router /admin [post]
func SaveAdmin(gin_context *gin.Context) {
	email := gin_context.Query("email")
	password := gin_context.Query("password")
	password = auth.HashPassword(password)
	admin, err := service.SaveAdmin(email, password)
	if err != nil {
		gin_context.JSON(http.StatusBadRequest, err)
		return
	}
	if admin == nil {
		gin_context.JSON(http.StatusBadRequest, "Admin already exists")
		return
	}
	// Create JSON with message and the token for the admin:
	token, err := auth.GenerateTokenFromMail(email)
	if err != nil {
		gin_context.JSON(http.StatusInternalServerError, "Something went wrong.")
		return
	}
	gin_context.JSON(http.StatusOK, gin.H{"message": "Admin saved", "token": token})
}

// @BasePath /api/v1

// GetAdmin godoc
// @Summary Endpoint used to get an admin
// @Schemes
// @Description gets an admin from the database
// @Tags admin
// @Accept json
// @Produce json
// @Param email query string true "Email of the admin" Format(email)
// @Param token header string true "Token of the admin" Format(token)
// @Success 200 {string} Admin found
// @Router /admin [get]
func GetAdmin(gin_context *gin.Context) {
	email := gin_context.Query("email")
	token := gin_context.GetHeader("token")
	if !verify_token(token) {
		gin_context.JSON(http.StatusBadRequest, gin.H{"message": "Token is not valid"})
		return
	}
	admin := service.GetAdmin(email)
	if admin == nil {
		gin_context.JSON(http.StatusNotFound, "Admin not found")
		return
	}
	gin_context.JSON(http.StatusOK, gin.H{"message": "Admin found", "email": admin.Email, "Time stamp": admin.TimeStamp})
}

// @BasePath /api/v1

// DeleteAdmin godoc
// @Summary Endpoint used to delete an admin
// @Schemes
// @Description deletes an admin from the database
// @Tags admin
// @Accept json
// @Produce json
// @Param email query string true "Email of the admin" Format(email)
// @Param token header string true "Token for authentification" Format(token)
// @Success 200 {string} Admin deleted
// @Router /admin [delete]
func DeleteAdmin(gin_context *gin.Context) {
	email := gin_context.Query("email")
	token := gin_context.GetHeader("token")
	if !verify_token(token) {
		gin_context.JSON(http.StatusBadRequest, gin.H{"message": "Token is not valid"})
		return
	}
	result, _ := service.DeleteAdmin(email)
	if result == "Admin not found" {
		gin_context.JSON(http.StatusBadRequest, result)
		return
	}
	gin_context.JSON(http.StatusOK, email+" is no longer an admin")
}

// @BasePath /api/v1

// LogIn godoc
// @Summary Endpoint used to log in an admin
// @Schemes
// @Description Given valid credentials, it returns a token
// @Tags admin
// @Accept json
// @Produce json
// @Param email query string true "Email of the admin" Format(email)
// @Param password query string true "Password of the admin" Format(password)
// @Success 200 {string} Log in succesful
// @Router /admin/login [post]
func LogIn(gin_context *gin.Context) {
	email := gin_context.Query("email")
	password := gin_context.Query("password")
	admin := service.GetAdmin(email)
	if admin == nil {
		gin_context.JSON(http.StatusNotFound, "Incorrect Credentials")
		return
	}
	if !auth.VerifyPassword(admin.Password, password) {
		gin_context.JSON(http.StatusBadRequest, "Incorrect Credentials")
		return
	}
	token, err := auth.GenerateTokenFromMail(email)
	if err != nil {
		gin_context.JSON(http.StatusInternalServerError, "Something went wrong.")
		return
	}
	gin_context.JSON(http.StatusOK, gin.H{"message": "Log In Succesful", "token": token})
}

func verify_token(token string) bool {
	_, err := auth.GetMailFromToken(token)
	return err == nil
}

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/", SaveAdmin)
			admin.POST("/login", LogIn)
			admin.GET("/", GetAdmin)
			admin.DELETE("/", DeleteAdmin)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	error := router.Run(":8080")
	if error != nil {
		panic(error)
	} // if we have an error we raise an exception
}
