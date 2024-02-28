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
		return nil, err
	}

	return admin.ToResponse(), nil
}
