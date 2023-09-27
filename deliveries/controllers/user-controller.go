package controllers

import (
	"sanberhub-test/entities/web"
	"sanberhub-test/services"

	"github.com/labstack/echo/v4"
)

type UserAndAccountControllers struct {
	usrAcc *services.UserAndAccountServices
}

func NewUserAndAccountControllers(userAccount *services.UserAndAccountServices) *UserAndAccountControllers {
	return &UserAndAccountControllers{
		usrAcc: userAccount,
	}
}

func (reg *UserAndAccountControllers) RegisterController(c echo.Context) error {
	var response web.Response
	var registerRequest web.RegisterRequest

	// binding json request from client
	err := c.Bind(&registerRequest)
	if err != nil {
		response.BadRequest("bad request", err.Error())
		return c.JSON(response.Code, response)
	}

	data, remark, message, err := reg.usrAcc.RegisterService(&registerRequest)
	if err != nil {
		response.InternalServerError("internal server error", err)
		return c.JSON(response.Code, response)
	}

	if message != "" {
		response.BadRequest(message, remark)
		return c.JSON(response.Code, response)
	}

	response.SuccessCreate("register successfully", data)
	return c.JSON(response.Code, response)
}

func (usr *UserAndAccountControllers) GetBalanceController(c echo.Context) error {
	var response web.Response

	accountNumber := c.Param("account_number")

	data, remark, message, err := usr.usrAcc.GetBalance(&accountNumber)
	if err != nil {
		response.InternalServerError("internal server error", err)
		return c.JSON(response.Code, response)
	}

	if message != "" {
		response.BadRequest(message, remark)
		return c.JSON(response.Code, response)
	}

	response.Success("get balance successfully", data)
	return c.JSON(response.Code, response)

}
