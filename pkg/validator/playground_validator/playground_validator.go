package playground_validator

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/hexley21/fixup/pkg/http/rest"
)

type playgroundValidator struct {
	validator *validator.Validate
}

func New() *playgroundValidator {
	validate := validator.New()
	err := validate.RegisterValidation("phone", phoneNumberValidator)
	if err != nil {
		log.Fatalf("failed to register phone validator: %v", err)
	}

	return &playgroundValidator{validator: validate}
}

func (v *playgroundValidator) Validate(i any) *rest.ErrorResponse {
	err := v.validator.Struct(i)
	if err != nil {
		return rest.NewInvalidArgumentsError(err)
	}

	return nil
}

func phoneNumberValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()

	re := regexp.MustCompile(`^[1-9]?[0-9]{7,14}$`)
	return re.MatchString(phone)
}
