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
	userService := services.NewUserAndAccountServices(userRepo, accountRepo)
	userAccountController := controllers.NewUserAndAccountControllers(userService)

	// deposit constructor
	depoService := services.NewDepoServices(accountRepo, mutationRepo)
	depoController := controllers.NewDepoControllers(depoService)

	// withdrawal constructor
	withdrawalService := services.NewWithdrawalServices(accountRepo, mutationRepo)
	withdrawalController := controllers.NewWithdrawalControllers(withdrawalService)

	// mutation constructor
	mutationService := services.NewMutationServices(mutationRepo)
	mutationController := controllers.NewMutationControllers(mutationService)

	v1 := "/api/v1"
	e.POST(v1+"/register", userAccountController.RegisterController)
	e.PUT(v1+"/deposit", depoController.DepositController)
	e.PUT(v1+"/withdrawal", withdrawalController.WithdrawalController)
	e.GET(v1+"/balance/:account_number", userAccountController.GetBalanceController)
	e.GET(v1+"/mutation/:account_number", mutationController.GetMutationsController)
}
