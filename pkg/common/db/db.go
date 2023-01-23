package db

import (
	"log"
	
	"gorm.io/driver/postgres"
	"github.com/ssjh23/PortfolioBackend/pkg/common/models"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.HeroSectionDescription{})

	return db
}
