package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/httpmessage"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
)

func (h Handler) userRegister(c echo.Context) error {
	var req param.RegisterRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	fieldError, err := h.userValidator.ValidateRegisterRequest(req)
	if err != nil {
		msg, code := httpmessage.ErrorCodeAndMsg(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldError,
		})
	}

	resp, err := h.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}
