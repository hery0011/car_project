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

func main() {
	apiAddress := fmt.Sprintf("%s:%d", config.AppConfiguration.API.Host, config.AppConfiguration.API.Port)
	if apiAddress == "" {
		log.Printf("EMPTY API_ADRESSSE: %s\n", apiAddress)
	}

	routes.GetRoutes(apiAddress)
}
