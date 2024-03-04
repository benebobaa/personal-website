package usecase

import (
	"github.com/go-playground/validator/v10"
	"log"
	"personal-website/delivery/http/dto/request"
	"personal-website/delivery/http/dto/response"
	"personal-website/domain/repository"
)

type AdminUsecase struct {
	Validate        *validator.Validate
	AdminRepository *repository.AdminRepository
}

func NewAdminUsecase(validator *validator.Validate, adminRepository *repository.AdminRepository) *AdminUsecase {
	return &AdminUsecase{
		Validate:        validator,
		AdminRepository: adminRepository,
	}
}

func (a *AdminUsecase) CreateAdmin(request *request.AdminRequest) (*response.AdminResponse, error) {

	err := a.Validate.Struct(request)
	if err != nil {
		log.Printf("error validating admin request:: %v", err)
		return nil, err
	}

	admin := request.ToEntity()

	err = a.AdminRepository.Create(admin)
	if err != nil {
		log.Printf("error creating admin:: %v", err)
		return nil, err
	}

	return admin.ToResponse(), nil
}

func (a *AdminUsecase) CreateHero(request *request.HeroRequest) error {

	err := a.Validate.Struct(request)
	if err != nil {
		log.Printf("error validating hero request:: %v", err)
		return err
	}

	err = a.AdminRepository.CreateHero(request.ToEntity())
	if err != nil {
		log.Printf("error creating hero:: %v", err)
		return err
	}

	return nil
}

func (a *AdminUsecase) UpdateHero(request *request.HeroUpdateRequest) error {

	err := a.Validate.Struct(request)
	if err != nil {
		log.Printf("error validating hero request:: %v", err)
		return err
	}

	err = a.AdminRepository.UpdateHero(request.ToEntity())

	if err != nil {
		log.Printf("error updating hero:: %v", err)
		return err
	}

	return nil
}

func (a *AdminUsecase) CreateSocmed(request *request.SocmedRequest) error {

	err := a.Validate.Struct(request)
	if err != nil {
		log.Printf("error validating social media request:: %v", err)
		return err
	}

	err = a.AdminRepository.CreateSocmed(request.ToEntity())
	if err != nil {
		log.Printf("error creating social media:: %v", err)
		return err
	}

	return nil
}

func (a *AdminUsecase) UpdateSocmed(request *request.SocmedUpdateRequest) error {

	err := a.Validate.Struct(request)
	if err != nil {
		log.Printf("error validating social media request:: %v", err)
		return err
	}

	err = a.AdminRepository.UpdateSocmed(request.ToEntity())

	if err != nil {
		log.Printf("error updating social media:: %v", err)
		return err
	}

	return nil
}
