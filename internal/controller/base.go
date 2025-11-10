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
	ListCategorie(*gin.Context)
	AjoutArticle(*gin.Context)
	DeleteArticle(*gin.Context)
	AjoutPanier(*gin.Context)
	DetailPanier(*gin.Context)
	DeletePanier(*gin.Context)
	AjoutCommande(*gin.Context)
	AssignCommande(*gin.Context)
	ListeCommandeOuvert(*gin.Context)
	ListeCommandeAssign(*gin.Context)
	ChercheCommercant(*gin.Context)
	ListCategories(*gin.Context)
	GetArticleDetail(*gin.Context)
	Checkout(*gin.Context)
	ListOrders(*gin.Context)
	GetTickets(*gin.Context)
	UpdateTicket(*gin.Context)
	AssignTicket(*gin.Context)
	FilterArticleByCommercant(*gin.Context)
	FilterArticleByName(*gin.Context)
	FilterArticleByCategorie(*gin.Context)
	ListLivreur(*gin.Context)
	GetLocation(*gin.Context)
	ListeArticleByCommercant(*gin.Context)
	FilterArticles(*gin.Context)
	UpdateArticle(*gin.Context)
	GetUserMenu(*gin.Context)
	GetMontantWallet(*gin.Context)
	RegisterCommercant(*gin.Context)
	RegisterLivreur(*gin.Context)
}

type livraisonHandler struct {
	db *gorm.DB
}

func NewHandler() LivraisonHandler {
	db := config.DatabaseConnex()

	return &livraisonHandler{db: db}
}
