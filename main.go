package main

import (
	"context"
	"log"
	"os"
	"sanberhub-test/config"
	"sanberhub-test/deliveries/route"
	"sanberhub-test/deliveries/validation"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	validation.AddCustomValidator()

	modelsDb := config.GetDbConnection()

	ctx := context.Background()
	e := echo.New()
	route.Build(e, modelsDb)
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))

	<-ctx.Done()
}
