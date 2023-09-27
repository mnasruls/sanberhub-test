package controllers

import (
	"sanberhub-test/entities/web"
	"sanberhub-test/services"

	"github.com/labstack/echo/v4"
)

type MutationControllers struct {
	mut *services.MutationServices
}

func NewMutationControllers(mutation *services.MutationServices) *MutationControllers {
	return &MutationControllers{
		mut: mutation,
	}
}

func (mtt *MutationControllers) GetMutationsController(c echo.Context) error {
	var response web.Response

	accountNumber := c.Param("account_number")

	data, remark, message, err := mtt.mut.GetUserMutations(&accountNumber)
	if err != nil {
		response.InternalServerError("internal server error", err)
		return c.JSON(response.Code, response)
	}

	if message != "" {
		response.BadRequest(message, remark)
		return c.JSON(response.Code, response)
	}

	response.Success("get mutations successfully", data)
	return c.JSON(response.Code, response)

}
