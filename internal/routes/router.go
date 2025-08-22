package routes

import (
	"car_project/internal/config"
	"car_project/internal/controller"
	"car_project/internal/jwt"
	"car_project/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRoutes(apiAddress string) {
	cHandler := controller.NewHandler()
	router := gin.Default()
	router.Use(cors.New(middleware.Cors()))

	// Route Swagger UI
	router.GET(config.SwaggerPath, ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		articleGroup := dashboardGroup.Group(config.ArticlePath, jwt.AuthMiddleware())
		{
			articleGroup.GET(config.ListArticle, cHandler.ListArticle)
			articleGroup.POST(config.AddArticle, cHandler.AjoutArticle)
			articleGroup.DELETE(config.DeleteArticle, cHandler.DeleteArticle)

			categorieGroup := articleGroup.Group(config.CategoriePath)
			{
				categorieGroup.GET(config.ListArticle, cHandler.ListCategorie)
			}

			panierGroup := articleGroup.Group(config.PanierPath)
			{
				panierGroup.POST(config.AjoutPanier, cHandler.AjoutPanier)
				panierGroup.GET(config.DetailPanier, cHandler.DetailPanier)
			}
		}
	}

	router.Run(apiAddress)
}
