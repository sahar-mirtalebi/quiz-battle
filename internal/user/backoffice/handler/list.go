package handler

import (
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/backoffice/dto"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/httpmessage"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func (h Handler) listAllUsers(c echo.Context) error {
	resp, err := h.backofficeUserSvc.ListAllUsers(c.Request().Context(), dto.ListAllUsersInput{})
	if err != nil {
		msg, code := httpmessage.ErrorCodeAndMsg(err)
		return echo.NewHTTPError(code, msg).SetInternal(richerror.New().WithError(err))
	}

	return c.JSON(http.StatusOK, resp.Users)
}
