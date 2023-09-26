package controllers

import (
	"sanberhub-test/entities/web"
	registService "sanberhub-test/services"

	"github.com/labstack/echo/v4"
)

type RegisterController struct {
	regist *registService.RegisterServices
}

func NewRegistController(register *registService.RegisterServices) *RegisterController {
	return &RegisterController{
		regist: register,
	}
}

func (reg *RegisterController) RegisterController(c echo.Context) error {
	var response web.Response
	var registerRequest web.RegisterRequest

	// binding json request from client
	err := c.Bind(&registerRequest)
	if err != nil {
		response.BadRequest("bad request", err.Error())
		return c.JSON(response.Code, response)
	}

	data, remark, message, err := reg.regist.RegisterService(&registerRequest)
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
