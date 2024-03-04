package router

import (
	"github.com/gofiber/fiber/v2"
	"personal-website/delivery/http/controller"
	"personal-website/domain"
)

type Router struct {
	adminController     controller.AdminController
	landingController   controller.LandingController
	dashboardController controller.DashboardController
}

func NewRouter(app *fiber.App, domain domain.Domain) {
	adminController := controller.NewAdminController(domain)

	landingController := controller.NewLandingController(domain)

	dashboardController := controller.NewDashboardController(domain)

	router := Router{
		adminController:     adminController,
		landingController:   landingController,
		dashboardController: dashboardController,
	}

	router.SetupRouter(app)
}

func (c *Router) SetupRouter(app *fiber.App) {
	app.Post("/admin", c.adminController.CreateAdmin)

	app.Get("/", c.landingController.Index)

	app.Post("/hero", c.adminController.CreateHero)
	app.Put("/hero/:id", c.adminController.UpdateHero)

	app.Post("/socmed", c.adminController.CreateSocmed)
	app.Patch("/socmed/:id", c.adminController.UpdateSocmed)

	// Dashboard Admin
	app.Get("/dashboard", c.dashboardController.Index)
}
