package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tomasff/brevis/internal/models"
)

type ShortService struct {
	Db models.Datastore
}

func (s *ShortService) GetShort(ctx *gin.Context) {

}

func (s *ShortService) CreateShort(ctx *gin.Context) {

}