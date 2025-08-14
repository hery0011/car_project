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
		authGroup.POST(config.Login, cHandler.Login)
		authGroup.POST(config.Logout, cHandler.Logout)
		authGroup.GET(config.Refresh, jwt.AuthMiddleware(), cHandler.Refresh)
	}

	userGroup := router.Group(config.UserPath, jwt.AuthMiddleware())
	{
		userGroup.POST(config.Creat, cHandler.CreatUser)
	}

	router.Run(apiAddress)
}
