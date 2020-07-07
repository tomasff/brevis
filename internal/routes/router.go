package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tomasff/brevis/internal/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/:id", controllers.GetShort)
	r.POST("/short", controllers.CreateShort)

	return r
}
