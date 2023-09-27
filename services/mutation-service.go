package services

import (
	"log"
	"sanberhub-test/deliveries/validation"
	"sanberhub-test/entities/web"
	"sanberhub-test/helpers"
	"sanberhub-test/repositories"

	"github.com/jinzhu/copier"
)

type MutationServices struct {
	mut *repositories.MutationRepositories
}

func NewMutationServices(mutation *repositories.MutationRepositories) *MutationServices {
	return &MutationServices{
		mut: mutation,
	}
}

func (mtt *MutationServices) GetUserMutations(accountNumber *string) (*[]web.MutationsResponse, interface{}, string, error) {

	// validation parameter
	getMutation := validation.Validator{
		UpdateBalanceRequest: &web.UpdateBalanceRequest{
			AccountNumber: *accountNumber,
		},
	}

	validate := getMutation.AccountNumberValidation()
	if validate != nil {
		log.Println("Error : ", helpers.JSONEncode(validate))
		return nil, validate, "bad request", nil
	}

	// get mutations data
	mutation, status, err := mtt.mut.GetMutation(accountNumber)
	if err != nil {
		log.Println("Error :", err)
		return nil, nil, "", err
	}

	if status != "" {
		log.Println("Error :", status)
		return nil, status, "bad request", nil
	}

	log.Println(helpers.JSONEncode(*mutation))

	// copy data from db to response
	var response []web.MutationsResponse

	err = copier.Copy(&response, mutation)
	if err != nil {
		log.Println("Error :", err)
		return nil, nil, "", err
	}

	return &response, nil, "", nil
}
