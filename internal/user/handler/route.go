package handler

import (
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetupRoutes(e *echo.Echo, presenceMiddleware echo.MiddlewareFunc) {
	userGroup := e.Group("/users")
	userGroup.POST("/login", h.userLogin)
	userGroup.POST("/register", h.userRegister)
	userGroup.GET("/profile", h.userProfile, echojwt.JWT(h.signKey), presenceMiddleware)
}
