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

func (r *AdminRepository) CreateHero(value *entity.Hero) error {
	result := r.db.Create(value)

	if result.Error != nil {

		log.Printf("error creating hero:: %v", result.Error)
		return result.Error
	}

	return nil
}

func (r *AdminRepository) UpdateHero(value *entity.Hero) error {
	result := r.db.Save(value)

	if result.Error != nil {

		log.Printf("error updating hero:: %v", result.Error)
		return result.Error
	}

	return nil
}

func (r *AdminRepository) CreateSocmed(value *entity.Socmed) error {
	result := r.db.Create(value)

	if result.Error != nil {

		log.Printf("error creating social media:: %v", result.Error)
		return result.Error
	}

	return nil
}

func (r *AdminRepository) UpdateSocmed(value *entity.Socmed) error {
	result := r.db.Model(value).Where("id = ?", value.ID).Updates(value)

	if result.Error != nil {

		log.Printf("error updating social media:: %v", result.Error)
		return result.Error
	}

	return nil
}
