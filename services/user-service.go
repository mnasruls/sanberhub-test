package services

import (
	"fmt"
	"log"
	"sanberhub-test/deliveries/validation"
	"sanberhub-test/entities/models"
	"sanberhub-test/entities/web"
	"sanberhub-test/helpers"
	"sanberhub-test/repositories"
	"time"

	"github.com/jinzhu/copier"
)

type UserAndAccountServices struct {
	usr *repositories.UserRepository
	acc *repositories.AccountRepositories
}

func NewUserAndAccountServices(user *repositories.UserRepository, account *repositories.AccountRepositories) *UserAndAccountServices {
	return &UserAndAccountServices{
		usr: user,
		acc: account,
	}
}

func (regService *UserAndAccountServices) RegisterService(req *web.RegisterRequest) (*web.RegistResponse, interface{}, string, error) {
	// validate client request
	registValidation := validation.Validator{
		RegisterRequest: req,
	}

	validate := registValidation.RegistValidation()
	if validate != nil {
		log.Println("Error : ", helpers.JSONEncode(validate))
		return nil, validate, "bad request", nil
	}

	// check user exist or not
	exist, err := regService.usr.CheckUser(&registValidation.NIK, &registValidation.NoHP)
	if err != nil {
		log.Println("Error : ", err)
		return nil, nil, "", err
	}

	if exist {
		log.Println("Error : nik or phone number already used")
		return nil, "nik or phone number already used", "bad request", nil
	}

	nowTime := time.Now().Format(helpers.DATE_STD_FORMAT)

	// copy request to models
	var usrModel models.User
	err = copier.Copy(&usrModel, req)
	if err != nil {
		log.Println("Error : ", err)
		return nil, nil, "", err
	}
	usrModel.CreatedAt = nowTime

	// create user
	id, err := regService.usr.CreateUser(&usrModel)
	if err != nil {
		log.Println("Error : ", err)
		return nil, nil, "", err
	}

	// generate account number user
	idStr := fmt.Sprintf("%04d", id)
	accountNumberStr := nowTime[:4] + req.NIK[:4] + idStr[len(idStr)-4:]

	accountUser := models.Account{
		UserId:        id,
		AccountNumber: accountNumberStr,
		CreatedAt:     nowTime,
		UpdatedAt:     nowTime,
	}

	// creates user account
	err = regService.acc.CreateAccount(&accountUser)
	if err != nil {
		log.Println("Error : ", err)
		return nil, nil, "", err
	}

	return &web.RegistResponse{
		AccountNumber: accountNumberStr,
	}, nil, "", nil

}

func (usr *UserAndAccountServices) GetBalance(accountNumber *string) (*web.UpdateBalanceResponse, interface{}, string, error) {

	// validation parameter
	getBalance := validation.Validator{
		UpdateBalanceRequest: &web.UpdateBalanceRequest{
			AccountNumber: *accountNumber,
		},
	}

	validate := getBalance.AccountNumberValidation()
	if validate != nil {
		log.Println("Error : ", helpers.JSONEncode(validate))
		return nil, validate, "bad request", nil
	}

	userAccount, status, err := usr.acc.GetAccount(accountNumber)
	if err != nil {
		log.Println("Error :", err)
		return nil, nil, "", err
	}

	if status != "" {
		log.Println("Error :", status)
		return nil, status, "bad request", nil
	}

	return &web.UpdateBalanceResponse{
		Balance: userAccount.Balance,
	}, nil, "", nil
}
