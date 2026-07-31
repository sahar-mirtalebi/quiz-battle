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

func (v Validator) ValidateRegisterRequest(ctx context.Context, req param.RegisterRequest) (map[string]string, error) {
	// TODO - we should verify phone number by verification code

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&req.Password, validation.Required, validation.Match(regexp.
			MustCompile(passwordRegex))),

		validation.Field(&req.PhoneNumber, validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex)).Error(errormessage.ErrorMsgPhoneNumberIsNotValid),
			validation.By(func(value interface{}) error {
				return v.checkPhoneNumberUniqueness(ctx, value)
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

func (v Validator) checkPhoneNumberUniqueness(ctx context.Context, value interface{}) error {
	phoneNumber := value.(string)
	if isUnique, err := v.Repo.IsPhoneNumberUnique(ctx, phoneNumber); err != nil || !isUnique {
		if err != nil {
			return err
		}

		if !isUnique {
			return fmt.Errorf("%s", errormessage.ErrorMsgPhoneNumberIsNotUnique)
		}
	}

	return nil
}
