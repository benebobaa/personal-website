package main

import (
	"embed"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	"log"
	httpLib "net/http"
	"personal-website/delivery/http"
	"personal-website/domain"
	"personal-website/shared/util"
	"personal-website/shared/util/token"
)

//go:embed resource/*
var resourcefs embed.FS

func main() {
	config, err := util.LoadConfig(".")
	validate := validator.New()

	tokenMaker, err := token.NewJWTMaker(config.TokenSymetricKey)
	if err != nil {
		log.Fatalf("Failed to create token maker: %v", err)
	}

	store := session.New()

	domain := domain.ConstructDomain(validate, tokenMaker, store, config)

	engine := html.NewFileSystem(httpLib.FS(resourcefs), ".html")
	app := http.NewHttpDelivery(domain, engine, config)
	err = app.Listen(fmt.Sprintf(":%s", config.PortApp))
	if err != nil {
		return
	}
}
