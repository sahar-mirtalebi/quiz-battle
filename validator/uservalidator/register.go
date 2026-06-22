package uservalidator

import (
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sahar-mirtalebi/quiz-battle/param"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/richerror"
)

func (v Validator) ValidateRegisterRequest(req param.RegisterRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateRegisterRequest"
	// TODO - we should verify phone number by verification code

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.
			MustCompile(passwordRegex))),

		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errormessage.ErrorMesagePhoneNumberIsNotvalid),
			validation.By(v.checkPhoneNumberUniqueness),
		),
	); err != nil {
		fieldError := map[string]string{}
		vErr, ok := err.(validation.Errors)
		if ok {
			for key, value := range vErr {
				if value != nil {
					fieldError[key] = value.Error()
				}
			}

		}
		return fieldError, richerror.New(op).WithError(err).
			WithMessage("invalid input").
			WithKind(richerror.KindInvalid).
			Withmeta(map[string]interface{}{
				"req": req,
			})
	}

	return nil, nil

}

func (v Validator) checkPhoneNumberUniqueness(value interface{}) error {
	phoneNumber := value.(string)
	if isUnique, err := v.Repo.IsPhoneNumberUnique(phoneNumber); err != nil || !isUnique {
		if err != nil {
			return err
		}

		if !isUnique {
			return fmt.Errorf(errormessage.ErrorMesagePhoneNumberIsNotUnique)
		}
	}

	return nil
}
