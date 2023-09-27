package validation

import (
	"sanberhub-test/entities/web"

	"github.com/thedevsaddam/govalidator"
)

type Validator struct {
	*web.RegisterRequest
	*web.UpdateBalanceRequest
}

func (r *Validator) RegistValidation() interface{} {
	// validation mandatory field
	validator := govalidator.New(govalidator.Options{
		Data: r,
		Rules: govalidator.MapData{
			"name":  []string{"required", "custom_char", "max:100"},
			"nik":   []string{"required", "numeric", "min:16"},
			"no_hp": []string{"required", "numeric", "between:6,20"},
		},
		RequiredDefault: true,
	}).ValidateStruct()

	if len(validator) > 0 {
		return validator
	}
	return nil
}

func (r *Validator) DepositValidation() interface{} {
	// validation mandatory field
	validator := govalidator.New(govalidator.Options{
		Data: r,
		Rules: govalidator.MapData{
			"account_number": []string{"required", "numeric", "max:12"},
			"nominal":        []string{"required", "numeric", "min:10000", "max:10000000"},
		},
		RequiredDefault: true,
	}).ValidateStruct()

	if len(validator) > 0 {
		return validator
	}
	return nil
}

func (r *Validator) WithdrawalValidation() interface{} {
	// validation mandatory field
	validator := govalidator.New(govalidator.Options{
		Data: r,
		Rules: govalidator.MapData{
			"account_number": []string{"required", "numeric", "max:12"},
			"nominal":        []string{"required", "numeric", "min:10000", "max:10000000"},
		},
		RequiredDefault: true,
	}).ValidateStruct()

	if len(validator) > 0 {
		return validator
	}
	return nil
}

func (r *Validator) AccountNumberValidation() interface{} {
	// validation mandatory field
	validator := govalidator.New(govalidator.Options{
		Data: r,
		Rules: govalidator.MapData{
			"account_number": []string{"required", "numeric", "max:12"},
		},
		RequiredDefault: true,
	}).ValidateStruct()

	if len(validator) > 0 {
		return validator
	}
	return nil
}
