package services

import (
	"log"
	"sanberhub-test/deliveries/validation"
	"sanberhub-test/entities/models"
	"sanberhub-test/entities/web"
	"sanberhub-test/helpers"
	"sanberhub-test/repositories"
	"time"
)

type DepoServices struct {
	acc *repositories.AccountRepositories
	mut *repositories.MutationRepositories
}

func NewDepoServices(account *repositories.AccountRepositories, mutation *repositories.MutationRepositories) *DepoServices {
	return &DepoServices{
		acc: account,
		mut: mutation,
	}
}

func (dep *DepoServices) DepositServices(req *web.UpdateBalanceRequest) (*web.UpdateBalanceResponse, interface{}, string, error) {
	// validate client request
	depoValidation := validation.Validator{
		UpdateBalanceRequest: req,
	}

	validate := depoValidation.DepositValidation()
	if validate != nil {
		log.Println("Error : ", helpers.JSONEncode(validate))
		return nil, validate, "bad request", nil
	}

	// check user account is exist or not
	userAccount, status, err := dep.acc.GetAccount(&req.AccountNumber)
	if err != nil {
		log.Println("Error :", err)
		return nil, nil, "", err
	}

	if status != "" {
		log.Println("Error :", status)
		return nil, status, "bad request", nil
	}

	// update balance to specific account number
	nowTime := time.Now().Format(helpers.DATE_STD_FORMAT)
	updatedBalance := userAccount.Balance + req.Nominal

	updatedUserAccount := models.Account{
		AccountNumber: req.AccountNumber,
		Balance:       updatedBalance,
		UpdatedAt:     nowTime,
	}

	err = dep.acc.UpdateBalance(&updatedUserAccount)
	if err != nil {
		log.Println("Error :", err)
		return nil, nil, "", err
	}

	// add mutation history
	mutation := models.Mutation{
		UserId:          userAccount.UserId,
		AccountNumber:   req.AccountNumber,
		KodeTransaksi:   helpers.DEPO_CODE,
		TransactionTime: nowTime,
		Nominal:         req.Nominal,
	}

	err = dep.mut.AddMutationHistory(&mutation)
	if err != nil {
		log.Println("Error :", err)
		return nil, nil, "", err
	}

	return &web.UpdateBalanceResponse{
		Balance: updatedBalance,
	}, nil, "", nil
}
