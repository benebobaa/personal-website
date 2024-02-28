package repository

import (
	"gorm.io/gorm"
	"log"
	"personal-website/domain/entity"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db}
}

func (r *AdminRepository) Create(value *entity.Admin) error {
	result := r.db.Create(value)

	if result.Error != nil {

		log.Printf("error creating admin:: %v", result.Error)
		return result.Error
	}

	return nil
}
