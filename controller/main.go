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
	test, err := auth.GetMailFromToken(token)
	if err != nil {
		gin_context.JSON(http.StatusBadRequest, err)
		return
	}
	if test != email {
		gin_context.JSON(http.StatusBadRequest, "Token doesn't match email")
		return
	}
	admin := service.GetAdmin(email)
	if admin == nil {
		gin_context.JSON(http.StatusNotFound, "Admin not found")
		return
	}
	gin_context.JSON(http.StatusOK, admin)
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
// @Success 200 {string} Admin deleted
// @Router /admin [delete]
func DeleteAdmin(gin_context *gin.Context) {
	email := gin_context.Query("email")
	result, _ := service.DeleteAdmin(email)
	if result == "Admin not found" {
		gin_context.JSON(http.StatusBadRequest, result)
		return
	}
	gin_context.JSON(http.StatusOK, email+" is no longer an admin")
}

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.POST("/", SaveAdmin)
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
