package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/middleware"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/httpmessage"
	"github.com/sahar-mirtalebi/quiz-battle/internal/user/dto"
)

func (h Handler) userProfile(c echo.Context) error {
	ctx := c.Request().Context()

	userID, err := middleware.UserID(c)
	if err != nil {
		return err
	}

	resp, err := h.userSvc.Profile(ctx, dto.ProfileInput{UserID: userID})
	if err != nil {
		msg, code := httpmessage.ErrorCodeAndMsg(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
