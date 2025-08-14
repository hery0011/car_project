package middleware

import (
	"car_project/internal/config"
	"log"

	"github.com/gin-contrib/cors"
)

func Cors() cors.Config {
	corsAllowOrigins := config.AppConfiguration.API.AllowOrigins
	if len(corsAllowOrigins) == 0 {
		log.Fatalln("CORS_ALLOW_ORIGINS n'est pas défini dans les variables d'environnement")
	}

	// Récupérer la valeur de la variable d'environnement CORS_ALLOW_ORIGINS
	config := cors.DefaultConfig()
	config.AllowOrigins = corsAllowOrigins
	config.AllowHeaders = []string{
		"Origin",
		"Content-Type",
		"Authorization",
		"Accept",
		"Accept-Language",
		"Content-Language",
		"Content-Length",
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
		"Access-Control-Allow-Origin",
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowCredentials = true
	return config
}
