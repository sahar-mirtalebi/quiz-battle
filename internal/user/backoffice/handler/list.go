package backofficeuserhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/httpmessage"
)

func (h Handler) listAllUsers(c echo.Context) error {
	resp, err := h.backofficeUserSvc.ListAllUsers(c.Request().Context())
	if err != nil {
		msg, code := httpmessage.ErrorCodeAndMsg(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
