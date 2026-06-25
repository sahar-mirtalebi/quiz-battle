package userhandler

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sahar-mirtalebi/quiz-battle/param"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/httpmessage"
)

func (h Handler) userProfile(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	userID := uint(claims["user_id"].(float64))

	resp, err := h.userSvc.Profile(param.ProfileRequest{UserID: userID})
	if err != nil {
		msg, code := httpmessage.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
