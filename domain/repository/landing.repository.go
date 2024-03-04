package repository

import (
	"gorm.io/gorm"
	"log"
	"personal-website/domain/entity"
)

type LandingRepository struct {
	db *gorm.DB
}

func NewLandingRepository(db *gorm.DB) *LandingRepository {
	return &LandingRepository{db}
}

func (r *LandingRepository) GetFirstHero() (*entity.Hero, error) {
	hero := &entity.Hero{ID: 1}

	result := r.db.First(hero)

	if result.Error != nil {
		log.Printf("error find hero:: %v", result.Error)
		return nil, result.Error
	}

	return hero, nil
}

func (r *LandingRepository) GetAllSocmeds() ([]entity.Socmed, error) {
	var socmeds []entity.Socmed

	result := r.db.Find(&socmeds)

	if result.Error != nil {
		log.Printf("error find social media:: %v", result.Error)
		return nil, result.Error
	}

	return socmeds, nil
}
