package controller

import (
	"github.com/gofiber/fiber/v2"
	"personal-website/domain"
)

type DashboardController interface {
	Index(ctx *fiber.Ctx) error
}

type dashboardControllerImpl struct {
	domain domain.Domain
}

func NewDashboardController(domain domain.Domain) DashboardController {
	return &dashboardControllerImpl{domain: domain}
}

func (d *dashboardControllerImpl) Index(ctx *fiber.Ctx) error {
	return ctx.Render("resource/views/admin/index", nil)
}
