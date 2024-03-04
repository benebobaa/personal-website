package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDatabaseConnection(dsn string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//if err := db.AutoMigrate(&entity.Admin{}, &entity.Hero{}, &entity.Socmed{}); err != nil {
	//	log.Fatal(err)
	//}

	return db
}
