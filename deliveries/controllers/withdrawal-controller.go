package controllers

import (
	"sanberhub-test/entities/web"
	"sanberhub-test/services"

	"github.com/labstack/echo/v4"
)

type WithdrawalControllers struct {
	with *services.WithdrawalServices
}

func NewWithdrawalControllers(withdrawal *services.WithdrawalServices) *WithdrawalControllers {
	return &WithdrawalControllers{
		with: withdrawal,
	}
}

func (wth *WithdrawalControllers) WithdrawalController(c echo.Context) error {
	var response web.Response
	var depoRequest web.UpdateBalanceRequest

	err := c.Bind(&depoRequest)
	if err != nil {
		response.BadRequest("bad request", err)
		return c.JSON(response.Code, response)
	}

	data, remark, status, err := wth.with.WithdrawalServices(&depoRequest)
	if err != nil {
		response.InternalServerError("internal server error", err)
		return c.JSON(response.Code, response)
	}

	if status != "" {
		response.BadRequest(status, remark)
		return c.JSON(response.Code, response)
	}

	response.Success("withdrawal successfully", data)
	return c.JSON(response.Code, response)
}
