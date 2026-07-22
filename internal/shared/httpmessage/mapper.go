// Package httpmessage maps application errors to HTTP responses.
package httpmessage

import (
	"errors"
	"net/http"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func ErrorCodeAndMsg(err error) (string, int) {
	var re richerror.RichError
	if errors.As(err, &re) {
		msg := re.Message()
		code := mapKindToHttpStatusCode(re.Kind())
		if code >= 500 {
			msg = errormessage.ErrorMsgSomethingWentWrong
		}
		return msg, code
	}

	return err.Error(), http.StatusBadRequest
}

func mapKindToHttpStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:

		return http.StatusUnprocessableEntity
	case richerror.KindForbiden:

		return http.StatusForbidden
	case richerror.KindNotFound:

		return http.StatusNotFound
	case richerror.KindUnexpected:

		return http.StatusInternalServerError
	default:

		return http.StatusBadRequest
	}
}
