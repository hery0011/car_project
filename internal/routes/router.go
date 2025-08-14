package routes

import (
	"car_project/internal/config"
	"car_project/internal/controller"
	"car_project/internal/jwt"
	"car_project/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRoutes(apiAddress string) {
	cHandler := controller.NewHandler()
	router := gin.Default()
	router.Use(cors.New(middleware.Cors()))

	authGroup := router.Group(config.AuthPath)
	{
		authGroup.POST(config.LoginPath, cHandler.Login)
		authGroup.POST(config.LogoutPath, cHandler.Logout)
		authGroup.GET(config.RefreshPath, jwt.AuthMiddleware(), cHandler.Refresh)
	}

	router.Run(apiAddress)
}
