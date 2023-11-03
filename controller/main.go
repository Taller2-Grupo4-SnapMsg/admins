package main

import (
	"admins/docs"
	"admins/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /admin/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @BasePath /api/v1

// SaveNumber godoc
// @Summary Test function for saving a number
// @Schemes
// @Description saves a number on a list
// @Tags admin
// @Accept json
// @Produce json
// @Param number query int true "Number to be saved" Format(int)
// @Success 200 {string} Number saved
// @Router /admin/numbers [post]
func SaveNumber(gin_context *gin.Context) {
	number, _ := strconv.Atoi(gin_context.Query("number"))
	service.SaveNumber(int32(number))
	gin_context.JSON(http.StatusOK, "Number saved")
}

// @BasePath /api/v1

// GetNumbers godoc
// @Summary Test function for getting a list of numbers
// @Schemes
// @Description gets a list of numbers
// @Tags admin
// @Accept json
// @Produce json
// @Success 200 {list} List of numbers
// @Router /admin/numbers [get]
func GetNumbers(gin_context *gin.Context) {
	gin_context.JSON(http.StatusOK, service.GetNumbers())
}

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/helloworld", Helloworld)
			admin.POST("/numbers", SaveNumber)
			admin.GET("/numbers", GetNumbers)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	error := router.Run(":8080")
	if error != nil {
		panic(error)
	} // if we have an error we raise an exception
}
