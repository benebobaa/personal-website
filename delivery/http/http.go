package http

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"personal-website/delivery/http/router"
	"personal-website/domain"
	"personal-website/shared/util"
	"time"
)

func NewHttpDelivery(domain domain.Domain, engine *html.Engine, c util.Config) *fiber.App {
	config := fiber.Config{
		AppName:           c.AppName,
		EnablePrintRoutes: true,
		JSONEncoder:       sonic.Marshal,
		JSONDecoder:       sonic.Unmarshal,
		Views:             engine,
	}

	if c.GOEnv == "production" {
		config.Prefork = true
	}

	app := fiber.New(config)
	app.Use(logger.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))
	app.Static("/static", "resource/public/web", fiber.Static{
		CacheDuration: 1 * time.Second,
	})

	router.NewRouter(app, domain)

	return app
}
