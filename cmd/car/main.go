package main

import (
	"car_project/internal/config"
	"car_project/internal/routes"
	"fmt"
	"log"
	"os"
)

var ROOT_FOLDER string

func init() {
	os.Setenv(config.ROOT_FOLDER_VAR, ROOT_FOLDER)
	config.Load()
}

// @title           Nom de votre API
// @version         1.0
// @description     Description de votre API
// @host            localhost:8082
// @BasePath        /
func main() {
	apiAddress := fmt.Sprintf("%s:%d", config.AppConfiguration.API.Host, config.AppConfiguration.API.Port)
	if apiAddress == "" {
		log.Printf("EMPTY API_ADRESSSE: %s\n", apiAddress)
	}

	routes.GetRoutes(apiAddress)
}
