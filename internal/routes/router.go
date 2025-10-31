package routes

import (
	"car_project/internal/config"
	"car_project/internal/controller"
	"car_project/internal/jwt"
	"car_project/internal/middleware"
	"car_project/internal/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRoutes(apiAddress string) {
	cHandler := controller.NewHandler()
	router := gin.Default()
	router.Use(cors.New(middleware.Cors()))

	router.Static("/uploads", "./uploads")

	// Route Swagger UI
	router.GET(config.SwaggerPath, ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	// --------------------------------------------------
	router.GET(config.ListCategories, cHandler.ListCategories)
	router.GET(config.GetArticleDetail, cHandler.GetArticleDetail)
	router.POST(config.Checkout, jwt.AuthMiddleware(), cHandler.Checkout)
	r := router.Group("/api/tickets")
	r.GET("", cHandler.GetTickets)
	// --------------------------------------------------

	// commercant articles filtering
	// --------------------------------------------------
	router.GET(config.FilterArticles, jwt.AuthMiddleware(), cHandler.FilterArticles)
	router.PUT(config.UpdateArticle, jwt.AuthMiddleware(), cHandler.UpdateArticle)
	// --------------------------------------------------

	routeDelivery := router.Group(config.Delivery, jwt.AuthMiddleware())
	{
		routeDelivery.GET(config.Tickets, cHandler.GetTickets)
		routeDelivery.PUT(config.UpdateTicket, cHandler.UpdateTicket) // pour éditer
		routeDelivery.PUT(config.AssignTicket, cHandler.AssignTicket) // pour assigner
	}

	router.GET(config.ListOrders, jwt.AuthMiddleware(), cHandler.ListOrders)

	authGroup := router.Group(config.AuthPath)
	{
		authGroup.POST(config.Login, cHandler.Login)
		authGroup.POST(config.Logout, cHandler.Logout)
		authGroup.GET(config.Refresh, jwt.AuthMiddleware(), cHandler.Refresh)
	}

	adminGroup := router.Group(config.AdminPath)
	{
		userGroup := adminGroup.Group(config.UserPath, jwt.AuthMiddleware())
		{
			userGroup.POST(config.Creat, cHandler.CreatUser)
			userGroup.DELETE(config.Delete, cHandler.DeleteUser)
			userGroup.PUT(config.Update, cHandler.UpdateUser)
		}

		profilGroup := adminGroup.Group(config.ProfilPath, jwt.AuthMiddleware())
		{
			profilGroup.GET(config.GetProfil, cHandler.GetListProfil)
			profilGroup.POST(config.AssignProfil, cHandler.AssignProfil)
		}
	}

	dashboardGroup := router.Group(config.DashPath)
	{
		articleGroup := dashboardGroup.Group(config.ArticlePath)
		{
			articleGroup.GET(config.ListArticle, cHandler.ListArticle)
			articleGroup.GET(config.ListeArticleByCommercant, cHandler.ListeArticleByCommercant)
			articleGroup.POST(config.AddArticle, jwt.AuthMiddleware(), cHandler.AjoutArticle)
			articleGroup.DELETE(config.DeleteArticle, cHandler.DeleteArticle)
			articleGroup.GET(config.FilterArticleByCommercant, cHandler.FilterArticleByCommercant)
			articleGroup.GET(config.FilterArticleByName, cHandler.FilterArticleByName)
			articleGroup.GET(config.FilterArticleByCategorie, cHandler.FilterArticleByCategorie)

			categorieGroup := articleGroup.Group(config.CategoriePath)
			{
				categorieGroup.GET(config.ListArticle, cHandler.ListCategorie)
			}

			panierGroup := articleGroup.Group(config.PanierPath)
			{
				panierGroup.POST(config.AjoutPanier, cHandler.AjoutPanier)
				panierGroup.GET(config.DetailPanier, cHandler.DetailPanier)
				panierGroup.DELETE(config.DeletePanier, cHandler.DeletePanier)
			}

			commandeGroup := articleGroup.Group(config.CommandePath, jwt.AuthMiddleware())
			{
				commandeGroup.POST(config.AjoutCommande, cHandler.AjoutCommande)
				commandeGroup.PUT(config.AssignCommande, cHandler.AssignCommande)
				commandeGroup.GET(config.ListCommandeCreer, cHandler.ListeCommandeOuvert)
				commandeGroup.GET(config.CommandeAssign, cHandler.ListeCommandeAssign)
			}

			commercantGroup := dashboardGroup.Group(config.CommercantPath, jwt.AuthMiddleware())
			{
				commercantGroup.POST(config.ChercheCommercant, cHandler.ChercheCommercant)
			}

		}

		livreurGroup := dashboardGroup.Group(config.LivreurPath, jwt.AuthMiddleware())
		{
			livreurGroup.POST(config.AjoutLivreur, cHandler.AjoutLivreur)
			livreurGroup.GET(config.ListLivreur, cHandler.ListLivreur)
			livreurGroup.GET(config.ListLivreurFilter, cHandler.GetLocation)

		}
	}

	// ✅ Nouvelle route WebSocket pour commerçants
	router.GET("/ws/commercant/:id", ws.HandleWS)

	router.Run(apiAddress)
}
