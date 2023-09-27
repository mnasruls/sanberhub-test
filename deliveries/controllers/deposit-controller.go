package controllers

import (
	"sanberhub-test/entities/web"
	"sanberhub-test/services"

	"github.com/labstack/echo/v4"
)

type DepoControllers struct {
	depo *services.DepoServices
}

func NewDepoControllers(deposit *services.DepoServices) *DepoControllers {
	return &DepoControllers{
		depo: deposit,
	}
}

func (dep *DepoControllers) DepositController(c echo.Context) error {
	var response web.Response
	var depoRequest web.UpdateBalanceRequest

	err := c.Bind(&depoRequest)
	if err != nil {
		response.BadRequest("bad request", err)
		return c.JSON(response.Code, response)
	}

	data, remark, status, err := dep.depo.DepositServices(&depoRequest)
	if err != nil {
		response.InternalServerError("internal server error", err)
		return c.JSON(response.Code, response)
	}

	if status != "" {
		response.BadRequest(status, remark)
		return c.JSON(response.Code, response)
	}

	response.Success("deposit successfully", data)
	return c.JSON(response.Code, response)
}
