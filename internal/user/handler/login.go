package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/httpmessage"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
)

func (h Handler) userLogin(c echo.Context) error {
	ctx := c.Request().Context()

	var req dto.LoginInput
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest).SetInternal(richerror.New().WithError(err).WithKind(richerror.KindInvalid))
	}

	fieldError, err := h.userValidator.ValidateLoginInput(ctx, req)
	if err != nil {
		msg, code := httpmessage.ErrorCodeAndMsg(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldError,
		})
	}

	resp, err := h.userSvc.Login(ctx, req)
	if err != nil {
		msg, code := httpmessage.ErrorCodeAndMsg(err)
		return echo.NewHTTPError(code, msg).SetInternal(richerror.New().WithError(err))
	}

	return c.JSON(http.StatusOK, resp)
}
