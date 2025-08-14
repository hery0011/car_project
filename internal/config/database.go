package config

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// DatabaseConnex : Fonction pour établir la connexion GORM à la base de données
func DatabaseConnex() *gorm.DB {
	once.Do(func() {
		dbDSN := AppConfiguration.DB_DSN
		if dbDSN == "" {
			log.Fatalln("Empty database configuration")
		}

		var err error
		// Ouvrir la connexion à la base de données avec GORM
		db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	})

	return db
}
