package httpmessage

import (
	"net/http"

	"github.com/sahar-mirtalebi/quiz-battle/pkg/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/richerror"
)

func Error(err error) (string, int) {
	switch err.(type) {
	case *richerror.RichError:
		re := err.(richerror.RichError)
		msg := re.Message()
		code := mapKindToHttpStatusCode(re.Kind())
		if code >= 500 {
			msg = errormessage.ErrorMsgSomeThingWentWrong
		}
		return msg, code
	default:

		return err.Error(), http.StatusBadRequest
	}
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
