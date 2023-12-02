package main

import (
	"admins/auth"
	"admins/docs"
	"admins/service"
	"admins/structs"
	"log"
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
// @Param credentials body string true "Email and password of the admin" SchemaExample({ "email": "admin@gmail.com", "password": "admin" })
// @Success 200 {string} Admin saved
// @Failure 400 {string} Admin already exists
// @Failure 400 {string} Invalid request data
// @Router /admin [post]
func SaveAdmin(gin_context *gin.Context) {
	var credentials structs.Credentials
	// Request a body with JSON:
	if err := gin_context.ShouldBindJSON(&credentials); err != nil {
		log.Fatal("[ERROR] Error binding JSON: ", err)
		gin_context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}
	email := credentials.Email
	password := credentials.Password
	password = auth.HashPassword(password)
	admin, err := service.SaveAdmin(email, password)
	if err != nil {
		gin_context.JSON(http.StatusBadRequest, err)
		return
	}
	if admin == nil {
		log.Println("[INFO] Admin ", email, " already exists")
		gin_context.JSON(http.StatusBadRequest, gin.H{"message": "Admin already exists"})
		return
	}
	// Create JSON with message and the token for the admin:
	token, err := auth.GenerateTokenFromMail(email)
	if err != nil {
		log.Fatal("[ERROR] Error generating token: ", err)
		gin_context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong."})
		return
	}
	log.Println("[INFO] Admin ", email, " has been registered")
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
// @Param token header string true "Token of the admin" Format(token)
// @Success 200 {string} Admin found
// @Failure 404 {string} Admin not found
// @Failure 400 {string} Token is not valid
// @Router /admin [get]
func GetAdmin(gin_context *gin.Context) {
	token := gin_context.GetHeader("token")
	email, err := auth.GetMailFromToken(token)
	if err != nil {
		gin_context.JSON(http.StatusBadRequest, gin.H{"message": "Token is not valid"})
		return
	}
	admin := service.GetAdmin(email)
	if admin == nil {
		gin_context.JSON(http.StatusNotFound, gin.H{"message": "Admin not found"})
		return
	}
	log.Println("[INFO] Admin ", email, " has been requested by token")
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
// @Failure 404 {string} Admin not found
// @Failure 400 {string} Token is not valid
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
// @Param loginRequest body string true "Email and password of the admin" SchemaExample({ "email": "admin@gmail.com", "password": "admin" })
// @Success 200 {string} Log in succesful
// @Failure 404 {string} Incorrect credentials
// @Failure 400 {string} Invalid request data
// @Failure 500 {string} Something went wrong
// @Router /admin/login [post]
func LogIn(gin_context *gin.Context) {
	var loginRequest structs.Credentials

	// Request a body with JSON:
	if err := gin_context.ShouldBindJSON(&loginRequest); err != nil {
		log.Fatal("[ERROR] Error binding JSON: ", err)
		gin_context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}
	email := loginRequest.Email
	password := loginRequest.Password
	admin := service.GetAdmin(email)
	if admin == nil {
		log.Fatal("[ERROR] Admin doesn't exist")
		gin_context.JSON(http.StatusNotFound, gin.H{"message": "Incorrect Credentials"})
		return
	}
	if !auth.VerifyPassword(admin.Password, password) {
		log.Fatal("[ERROR] Password doesn't match")
		gin_context.JSON(http.StatusNotFound, gin.H{"message": "Incorrect Credentials"})
		return
	}
	token, err := auth.GenerateTokenFromMail(email)
	if err != nil {
		log.Fatal("[ERROR] Error generating token: ", err)
		gin_context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong."})
		return
	}
	log.Println("[INFO] Admin ", email, " has logged in")
	gin_context.JSON(http.StatusOK, gin.H{"message": "Log In Succesful", "token": token})
}

// @BasePath /api/v1

// GetHealth godoc
// @Summary Endpoint used to check if the server is running
// @Schemes
// @Description returns a message if the server is running
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {string} Server is running
// @Router /admin/health [get]
func GetHealth(gin_context *gin.Context) {
	// Here we can add a more complex response if needed.
	gin_context.JSON(http.StatusOK, gin.H{
		"status":        "ok",
		"description":   "Micro service that handles admin authentication.",
		"creation_date": "30-10-2023",
	})
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
			admin.GET("/health", GetHealth)
			admin.DELETE("/", DeleteAdmin)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	error := router.Run(":8080")
	if error != nil {
		panic(error)
	} // if we have an error we raise an exception
}
