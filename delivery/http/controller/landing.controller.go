package controller

import (
	"github.com/gofiber/fiber/v2"
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
	return ctx.Render("resource/views/landing/index", nil)
}
