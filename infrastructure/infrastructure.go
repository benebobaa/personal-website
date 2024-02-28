package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"personal-website/domain/entity"
)

func NewDatabaseConnection(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&entity.Admin{}); err != nil {
		log.Fatal(err)
	}

	return db
}
