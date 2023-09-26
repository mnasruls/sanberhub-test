package validation

import (
	"sanberhub-test/entities/web"

	"github.com/thedevsaddam/govalidator"
)

type Validator struct {
	*web.RegisterRequest
	*web.DepositRequest
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
