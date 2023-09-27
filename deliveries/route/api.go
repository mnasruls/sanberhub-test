package route

import (
	"sanberhub-test/deliveries/controllers"
	"sanberhub-test/repositories"
	"sanberhub-test/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Build(e *echo.Echo, db *gorm.DB) {

	// repository constructor inject
	userRepo := repositories.NewUserRepository(db)
	accountRepo := repositories.NewAccountRepositories(db)
	mutationRepo := repositories.NewMutationRepositories(db)

	// register constructor
	regiterService := services.NewRegisterServices(userRepo, accountRepo)
	registerController := controllers.NewRegistController(regiterService)

	depoService := services.NewDepoServices(accountRepo, mutationRepo)
	depoController := controllers.NewDepoControllers(depoService)

	v1 := "/api/v1"
	e.POST(v1+"/register", registerController.RegisterController)
	e.PUT(v1+"/deposit", depoController.DepositController)
}
