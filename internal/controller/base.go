package controller

import (
	"car_project/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LivraisonHandler interface {
	Login(*gin.Context)
	Logout(*gin.Context)
	Refresh(*gin.Context)
}

type livraisonHandler struct {
	db *gorm.DB
}

func NewHandler() LivraisonHandler {
	db := config.DatabaseConnex()

	return &livraisonHandler{db: db}
}
