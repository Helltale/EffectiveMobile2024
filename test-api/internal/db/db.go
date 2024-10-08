// internal/db/database.go
package db

import (
	"os"

	"github.com/helltale/EffectiveMobile2024/test-api/internal/logger"
	"github.com/helltale/EffectiveMobile2024/test-api/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Fatal("err: can not get .env")
	}

	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	link := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"

	DB, err = gorm.Open("postgres", link)
	if err != nil {
		logger.Log.Fatal("err: can not connect with db", err)
	}

	logger.Log.Info("connected to the database")
	DB.AutoMigrate(&models.Song{})
}
