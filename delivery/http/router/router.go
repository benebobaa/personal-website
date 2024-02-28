package router

import (
	"github.com/gofiber/fiber/v2"
	"personal-website/delivery/http/controller"
	"personal-website/domain"
)

type Router struct {
	adminController   controller.AdminController
	landingController controller.LandingController
}

func NewRouter(app *fiber.App, domain domain.Domain) {
	adminController := controller.NewAdminController(domain)

	landingController := controller.NewLandingController(domain)

	router := Router{
		adminController:   adminController,
		landingController: landingController,
	}

	router.SetupRouter(app)
}

func (c *Router) SetupRouter(app *fiber.App) {
	app.Post("/admin", c.adminController.CreateAdmin)

	app.Get("/", c.landingController.Index)
}
