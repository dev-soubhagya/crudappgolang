package routes

import (
	"crudappgolang/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/upload", controllers.UploadExcel)
	r.GET("/data", controllers.ViewData)
	r.PUT("/edit/:id", controllers.EditRecord)
	return r
}
