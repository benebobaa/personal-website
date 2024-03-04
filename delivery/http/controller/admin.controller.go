package controller

import (
	"github.com/gofiber/fiber/v2"
	"personal-website/delivery/http/dto/request"
	"personal-website/domain"
	"strconv"
)

type AdminController interface {
	CreateAdmin(ctx *fiber.Ctx) error
	CreateHero(ctx *fiber.Ctx) error
	UpdateHero(ctx *fiber.Ctx) error
	CreateSocmed(ctx *fiber.Ctx) error
	UpdateSocmed(ctx *fiber.Ctx) error
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

func (a *adminControllerImpl) CreateHero(ctx *fiber.Ctx) error {
	requestBody := new(request.HeroRequest)

	if err := ctx.BodyParser(requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := a.domain.AdminUsecase.CreateHero(requestBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Hero created successfully",
	})
}

func (a *adminControllerImpl) UpdateHero(ctx *fiber.Ctx) error {
	requestBody := new(request.HeroUpdateRequest)
	id := ctx.Params("id", "")

	if err := ctx.BodyParser(requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	heroId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid hero ID",
		})
	}

	requestBody.ID = heroId

	err = a.domain.AdminUsecase.UpdateHero(requestBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hero updated successfully",
	})
}

func (a *adminControllerImpl) CreateSocmed(ctx *fiber.Ctx) error {
	requestBody := new(request.SocmedRequest)

	if err := ctx.BodyParser(requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := a.domain.AdminUsecase.CreateSocmed(requestBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Social media created successfully",
	})
}

func (a *adminControllerImpl) UpdateSocmed(ctx *fiber.Ctx) error {
	requestBody := new(request.SocmedUpdateRequest)
	id := ctx.Params("id", "")

	if err := ctx.BodyParser(requestBody); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	heroId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid hero ID",
		})
	}

	requestBody.ID = heroId

	err = a.domain.AdminUsecase.UpdateSocmed(requestBody)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Social media updated successfully",
	})
}
