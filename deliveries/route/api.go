package route

import (
	"sanberhub-test/deliveries/controllers"
	"sanberhub-test/repositories"
	"sanberhub-test/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Build(e *echo.Echo, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	accountRepo := repositories.NewAccountRepositories(db)
	regiterService := services.NewRegisterServices(userRepo, accountRepo)
	registerController := controllers.NewRegistController(regiterService)

	v1 := "/api/v1"
	e.POST(v1+"/register", registerController.RegisterController)
}
