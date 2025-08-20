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
	CreatUser(*gin.Context)
	DeleteUser(*gin.Context)
	UpdateUser(*gin.Context)
	GetListProfil(*gin.Context)
	AssignProfil(*gin.Context)
	ListArticle(*gin.Context)
}

type livraisonHandler struct {
	db *gorm.DB
}

func NewHandler() LivraisonHandler {
	db := config.DatabaseConnex()

	return &livraisonHandler{db: db}
}
