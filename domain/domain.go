package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/session"
	"personal-website/domain/repository"
	"personal-website/domain/usecase"
	"personal-website/infrastructure"
	"personal-website/shared/util"
	"personal-website/shared/util/token"
)

type Domain struct {
	Validate   *validator.Validate
	TokenMaker token.Maker
	Store      *session.Store

	AdminUsecase   *usecase.AdminUsecase
	LandingUsecase *usecase.LandingUsecase
}

func ConstructDomain(
	validate *validator.Validate,
	tokenMaker token.Maker,
	store *session.Store,
	c util.Config,
) Domain {
	databaseConn := infrastructure.NewDatabaseConnection(c.DBDsn)

	//repository
	adminRepository := repository.NewAdminRepository(databaseConn)
	landingRepository := repository.NewLandingRepository(databaseConn)

	adminUsecase := usecase.NewAdminUsecase(validate, adminRepository)
	landingUsecase := usecase.NewLandingUsecase(validate, landingRepository)

	return Domain{
		Validate:       validate,
		TokenMaker:     tokenMaker,
		Store:          store,
		AdminUsecase:   adminUsecase,
		LandingUsecase: landingUsecase,
	}
}
