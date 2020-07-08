package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tomasff/brevis/internal/models"
)

func SetupRouter(db *models.DB) *gin.Engine {
	r := gin.Default()

	shortService := &ShortService{db}

	r.GET("/:id", shortService.GetShort)
	r.POST("/short", shortService.CreateShort)

	return r
}
