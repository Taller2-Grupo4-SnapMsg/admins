package main

import (
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
	admin, error := service.SaveAdmin(email, password)
	if error != nil {
		gin_context.JSON(http.StatusBadRequest, error)
		return
	}
	gin_context.JSON(http.StatusOK, admin.Email+" is now an admin")
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
// @Success 200 {string} Admin found
// @Router /admin [get]
func GetAdmin(gin_context *gin.Context) {
	email := gin_context.Query("email")
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
