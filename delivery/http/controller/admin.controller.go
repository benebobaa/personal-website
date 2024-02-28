package controller

import (
	"github.com/gofiber/fiber/v2"
	"personal-website/delivery/http/dto/request"
	"personal-website/domain"
)

type AdminController interface {
	CreateAdmin(ctx *fiber.Ctx) error
}

type adminControllerImpl struct {
	domain domain.Domain
}

func NewAdminController(domain domain.Domain) AdminController {
	return &adminControllerImpl{
		domain: domain,
	}
}

func (a *adminControllerImpl) CreateAdmin(ctx *fiber.Ctx) error {
	requestBody := new(request.AdminRequest)

	if err := ctx.BodyParser(requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	response, err := a.domain.AdminUsecase.CreateAdmin(requestBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
