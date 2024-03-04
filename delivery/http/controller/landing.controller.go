package controller

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"personal-website/domain"
)

type LandingController interface {
	Index(ctx *fiber.Ctx) error
}

type landingControllerImpl struct {
	domain domain.Domain
}

func NewLandingController(domain domain.Domain) LandingController {
	return &landingControllerImpl{domain: domain}
}

func (l *landingControllerImpl) Index(ctx *fiber.Ctx) error {

	hero, err := l.domain.LandingUsecase.GetHero()
	if err != nil {
		log.Printf("error getting hero:: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	socmeds, err := l.domain.LandingUsecase.GetSocmed()
	if err != nil {
		log.Printf("error getting socmeds:: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Render("resource/views/landing/index", fiber.Map{
		"Hero":    hero,
		"Socmeds": socmeds,
	})
}
