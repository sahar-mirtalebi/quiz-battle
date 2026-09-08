package httpmessage

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/errormessage"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func TestWrappedUnexpectedErrorHTTPResponse(t *testing.T) {
	cause := errors.New("private database details")
	err := richerror.New().WithError(richerror.New().WithError(cause).WithKind(richerror.KindUnexpected))
	msg, code := ErrorCodeAndMsg(err)
	if code != http.StatusInternalServerError || msg != errormessage.ErrorMsgSomethingWentWrong {
		t.Fatalf("unexpected public error: %d %q", code, msg)
	}
	httpErr := echo.NewHTTPError(code, msg).SetInternal(err)
	if !errors.Is(httpErr, cause) {
		t.Fatal("HTTP error lost the original cause")
	}
	e := echo.New()
	recorder := httptest.NewRecorder()
	e.DefaultHTTPErrorHandler(httpErr, e.NewContext(httptest.NewRequest(http.MethodGet, "/", nil), recorder))
	if recorder.Code != code || recorder.Body.String() != "{\"message\":\""+msg+"\"}\n" {
		t.Fatalf("unexpected Echo response: %d %s", recorder.Code, recorder.Body.String())
	}
}
