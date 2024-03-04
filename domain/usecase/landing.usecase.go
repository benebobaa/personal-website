package usecase

import (
	"github.com/go-playground/validator/v10"
	"log"
	"personal-website/delivery/http/dto/response"
	"personal-website/domain/entity"
	"personal-website/domain/repository"
)

type LandingUsecase struct {
	Validate          *validator.Validate
	LandingRepository *repository.LandingRepository
}

func NewLandingUsecase(validate *validator.Validate, landingRepository *repository.LandingRepository) *LandingUsecase {
	return &LandingUsecase{Validate: validate,
		LandingRepository: landingRepository,
	}
}

func (l *LandingUsecase) GetHero() (*response.HeroResponse, error) {
	hero, err := l.LandingRepository.GetFirstHero()

	if err != nil {
		log.Printf("error getting hero:: %v", err)
		return nil, err
	}

	return hero.ToResponse(), nil
}

func (l *LandingUsecase) GetSocmed() ([]response.SocmedResponse, error) {
	socmeds, err := l.LandingRepository.GetAllSocmeds()

	if err != nil {
		log.Printf("error getting social media:: %v", err)
		return nil, err
	}

	return entity.ToSocmedResponses(socmeds), nil
}
