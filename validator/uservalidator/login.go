package uservalidator

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sahar-mirtalebi/quiz-battle/param"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/richerror"
)

func (v Validator) ValidateLoginRequest(req param.LoginRequest) (map[string]string, error) {
	const op = "uservalidator.ValidateLoginRequest"
	// TODO - we should verify phone number by verification code
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Password, validation.Required),
		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errormessage.ErrorMesagePhoneNumberIsNotvalid),
			validation.By(v.doesPhoneNumberExist),
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

func (v Validator) doesPhoneNumberExist(value interface{}) error {
	phoneNumber := value.(string)
	_, err := v.Repo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		//fmt.Errorf(errormessage.ErrorMsgNotFound)
	}

	return nil
}
