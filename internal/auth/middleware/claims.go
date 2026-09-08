package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
)

func UserID(c echo.Context) (uint, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok || !token.Valid {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "invalid token").SetInternal(richerror.New().WithMessage("invalid token").WithKind(richerror.KindForbiden))
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims").SetInternal(richerror.New().WithMessage("invalid token claims").WithKind(richerror.KindForbiden))
	}

	userID, ok := claims["user_id"].(float64)
	if !ok || userID <= 0 {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "invalid user_id").SetInternal(richerror.New().WithMessage("invalid user_id").WithKind(richerror.KindForbiden))
	}

	return uint(userID), nil
}
