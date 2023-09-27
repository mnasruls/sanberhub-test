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

type WithdrawalServices struct {
	acc *repositories.AccountRepositories
	mut *repositories.MutationRepositories
}

func NewWithdrawalServices(account *repositories.AccountRepositories, mutation *repositories.MutationRepositories) *WithdrawalServices {
	return &WithdrawalServices{
		acc: account,
		mut: mutation,
	}
}

func (with *WithdrawalServices) WithdrawalServices(req *web.UpdateBalanceRequest) (*web.UpdateBalanceResponse, interface{}, string, error) {
	// validate client request
	withdrawalValidation := validation.Validator{
		UpdateBalanceRequest: req,
	}

	validate := withdrawalValidation.WithdrawalValidation()
	if validate != nil {
		log.Println("Error : ", helpers.JSONEncode(validate))
		return nil, validate, "bad request", nil
	}

	// check user account is exist or not
	userAccount, status, err := with.acc.GetAccount(&req.AccountNumber)
	if err != nil {
		log.Println("Error :", err)
		return nil, nil, "", err
	}

	if status != "" {
		log.Println("Error :", status)
		return nil, status, "bad request", nil
	}

	// check if user balance is enough to witdrawal or no
	if req.Nominal > userAccount.Balance {
		status = "balance not enough"
		log.Println("Error :", status)
		return nil, status, "bad request", nil
	}

	// update balance to specific account number
	nowTime := time.Now().Format(helpers.DATE_STD_FORMAT)
	updatedBalance := userAccount.Balance - req.Nominal

	updatedUserAccount := models.Account{
		AccountNumber: req.AccountNumber,
		Balance:       updatedBalance,
		UpdatedAt:     nowTime,
	}

	err = with.acc.UpdateBalance(&updatedUserAccount)
	if err != nil {
		log.Println("Error :", err)
		return nil, nil, "", err
	}

	// add mutation history
	mutation := models.Mutation{
		UserId:          userAccount.UserId,
		AccountNumber:   req.AccountNumber,
		KodeTransaksi:   helpers.WITHDRAWAL_CODE,
		TransactionTime: nowTime,
		Nominal:         -req.Nominal,
	}

	err = with.mut.AddMutationHistory(&mutation)
	if err != nil {
		log.Println("Error :", err)
		return nil, nil, "", err
	}

	return &web.UpdateBalanceResponse{
		Balance: updatedBalance,
	}, nil, "", nil

}
