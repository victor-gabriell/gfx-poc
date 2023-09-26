package db

import (
	"fmt"
	"log"

	"github.com/victor-gabriell/gfx/internal/company/model"
	"github.com/victor-gabriell/gfx/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config *config.ApplicationConfig) *gorm.DB {

	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.DB,
		config.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	return db
}

func CreateEntities(db *gorm.DB) {
	err := db.AutoMigrate(&model.Company{})
	if err != nil {
		log.Fatalf("error while auto migrate %+v", err)
	}
}
