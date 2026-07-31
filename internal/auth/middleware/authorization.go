// Package middleware contains authentication and authorization middleware.
package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	authdomain "github.com/sahar-mirtalebi/quiz-battle/internal/auth/domain"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type AuthorizationService interface {
	CheckAccess(ctx context.Context, userID uint, role userdomain.Role, permissions []authdomain.PermissionTitle) (bool, error)
}

func Authorization(authzSvc AuthorizationService, permissions ...authdomain.PermissionTitle) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims")
			}

			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing user_id")
			}

			roleFloat, ok := claims["role"].(float64)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing role")
			}

			allowed, err := authzSvc.CheckAccess(c.Request().Context(), uint(userIDFloat), userdomain.Role(roleFloat), permissions)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError)
			}

			if !allowed {
				return echo.NewHTTPError(http.StatusForbidden, "user not allowed")
			}

			return next(c)
		}
	}
}
