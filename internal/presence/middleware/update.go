// Package middleware contains HTTP middleware for tracking user presence.
package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	authmiddleware "github.com/sahar-mirtalebi/quiz-battle/internal/auth/middleware"
	"github.com/sahar-mirtalebi/quiz-battle/internal/presence/dto"
	"github.com/sahar-mirtalebi/quiz-battle/internal/shared/timestamp"
)

type PresenceService interface {
	Upsert(ctx context.Context, input dto.UpsertPresenceInput) (dto.UpsertPresenceOutput, error)
}

func Upsert(service PresenceService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID, err := authmiddleware.UserID(c)
			if err != nil {
				return err
			}

			_, err = service.Upsert(c.Request().Context(), dto.UpsertPresenceInput{
				UserID:    userID,
				Timestamp: timestamp.Now(),
			})
			if err != nil {
				c.Logger().Errorf("update user presence: %v", err)
			}

			return next(c)
		}
	}
}
