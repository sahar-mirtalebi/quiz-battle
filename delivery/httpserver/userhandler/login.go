package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sahar-mirtalebi/quiz-battle/param"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/httpmessage"
)

func (h Handler) userLogin(c echo.Context) error {
	var req param.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	fieldError, err := h.userValidator.ValidateLoginRequest(req)
	if err != nil {
		msg, code := httpmessage.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldError,
		})
	}

	resp, err := h.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
