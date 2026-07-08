package backofficeuserhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/httpmessage"
)

func (h Handler) listAllUsers(c echo.Context) error {
	resp, err := h.backofficeUserSvc.ListAllUser()
	if err != nil {
		msg, code := httpmessage.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
