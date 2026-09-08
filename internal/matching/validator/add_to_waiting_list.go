package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sahar-mirtalebi/quiz-battle/internal/matching/dto"
	quizdomain "github.com/sahar-mirtalebi/quiz-battle/internal/quiz/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func (v Validator) ValidateAddToWaitingListInput(req dto.AddToWaitingListInput) (map[string]string, error) {
	// TODO - we should verify phone number by verification code

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.Category, validation.Required,
			validation.By(v.isCategoryValid),
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

func (v Validator) isCategoryValid(value interface{}) error {
	category := value.(quizdomain.Category)

	if !category.IsValid() {
		return richerror.New().WithMessage(errormessage.ErrorMsgCategoryIsNotValid).WithKind(richerror.KindInvalid)
	}

	return nil
}
