package services

import (
	"fmt"
	"log"
	"sanberhub-test/deliveries/validation"
	"sanberhub-test/entities/models"
	"sanberhub-test/entities/web"
	"sanberhub-test/helpers"
	"sanberhub-test/repositories"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
)

type RegisterServices struct {
	usr *repositories.UserRepository
	acc *repositories.AccountRepositories
}

func NewRegisterServices(user *repositories.UserRepository, account *repositories.AccountRepositories) *RegisterServices {
	return &RegisterServices{
		usr: user,
		acc: account,
	}
}

func (regService *RegisterServices) RegisterService(req *web.RegisterRequest) (*web.RegistResponse, interface{}, string, error) {
	// validate client request
	registValidation := validation.Validator{
		RegisterRequest: req,
	}

	validate := registValidation.RegistValidation()
	if validate != nil {
		log.Println("Error : ", helpers.JSONEncode(validate))
		return nil, validate, "invalid validation", nil
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
	accountNumber, err := strconv.Atoi(accountNumberStr)
	if err != nil {
		log.Println("Error : ", err)
		return nil, nil, "", err
	}

	accountUser := models.Account{
		UserId:        id,
		AccountNumber: accountNumber,
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
		AccountNumber: accountNumber,
	}, nil, "", nil

}
