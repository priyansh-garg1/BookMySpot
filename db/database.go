package db

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/priyansh-garg1/ticket-verse/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig) *gorm.DB {
	uri := fmt.Sprintf(`
	    host=%s user=%s dbname=%s password=%s sslmode=%s port=5432`,
		config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBSSLMode
	)

	db,err := gorm.Open(postgres.Open(uri), &&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err!= nil {
        log.Fatalf("Failed to connect to the database %e",err)
    }

	log.Info("Connected to database")

	return db
}

