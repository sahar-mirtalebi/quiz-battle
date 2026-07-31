package uservalidator

import (
	"context"
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
)

func (v Validator) ValidateLoginRequest(ctx context.Context, req param.LoginRequest) (map[string]string, error) {
	// TODO - we should verify phone number by verification code
	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Password, validation.Required),
		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errormessage.ErrorMsgPhoneNumberIsNotValid),
			validation.By(func(value interface{}) error {
				return v.doesPhoneNumberExist(ctx, value)
			}),
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
		return fieldError, richerror.New().WithError(err).
			WithMessage("invalid input").
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{
				"req": req,
			})
	}

	return nil, nil

}

func (v Validator) doesPhoneNumberExist(ctx context.Context, value interface{}) error {
	phoneNumber := value.(string)
	_, err := v.Repo.GetUserByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return fmt.Errorf("%s", errormessage.ErrorMsgNotFound)
	}

	return nil
}
