package userhandler

import (
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetUpUserRoute(e *echo.Echo) {
	userGroup := e.Group("/users")
	userGroup.POST("/login", h.userLogin)
	userGroup.POST("/register", h.userRegister)
	userGroup.GET("/profile", h.userProfile, echojwt.JWT(h.signKey))
}
