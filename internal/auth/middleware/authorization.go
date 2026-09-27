// Package middleware contains authentication and authorization middleware.
package middleware

import (
	"context"
	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/dto"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	authdomain "github.com/sahar-mirtalebi/quiz-battle/internal/auth/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/richerror"
	userdomain "github.com/sahar-mirtalebi/quiz-battle/internal/user/domain"
)

type AuthorizationService interface {
	CheckAccess(ctx context.Context, input dto.CheckAccessInput) (dto.CheckAccessOutput, error)
}

func Authorization(authzSvc AuthorizationService, permissions ...authdomain.PermissionTitle) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing token").SetInternal(richerror.New().WithMessage("missing token").WithKind(richerror.KindForbiden))
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token claims").SetInternal(richerror.New().WithMessage("invalid token claims").WithKind(richerror.KindForbiden))
			}

			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing user_id").SetInternal(richerror.New().WithMessage("missing user_id").WithKind(richerror.KindForbiden))
			}

			roleFloat, ok := claims["role"].(float64)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing role").SetInternal(richerror.New().WithMessage("missing role").WithKind(richerror.KindForbiden))
			}

			allowed, err := authzSvc.CheckAccess(c.Request().Context(), dto.CheckAccessInput{UserID: uint(userIDFloat), Role: userdomain.Role(roleFloat), Permissions: permissions})
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(richerror.New().WithError(err).WithKind(richerror.KindUnexpected))
			}

			if !allowed.Allowed {
				return echo.NewHTTPError(http.StatusForbidden, "user not allowed").SetInternal(richerror.New().WithMessage("user not allowed").WithKind(richerror.KindForbiden))
			}

			return next(c)
		}
	}
}
