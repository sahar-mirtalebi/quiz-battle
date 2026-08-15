package handler

import (
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	authdomain "github.com/sahar-mirtalebi/quiz-battle/internal/auth/domain"
	"github.com/sahar-mirtalebi/quiz-battle/internal/auth/middleware"
)

func (h Handler) SetupRoutes(e *echo.Echo) {
	backofficeGroup := e.Group("/backoffice/users")
	backofficeGroup.GET("", h.listAllUsers, echojwt.JWT(h.signKey), middleware.Authorization(h.authzSvc, authdomain.UserListPermission))
}
