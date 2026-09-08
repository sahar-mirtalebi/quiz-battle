package handler

import (
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetupRoutes(e *echo.Echo, presenceMiddleware echo.MiddlewareFunc) {
	userGroup := e.Group("/matching")
	userGroup.POST("/add-to-waiting-list", h.addToWaitingList, echojwt.JWT(h.signKey), presenceMiddleware)

}
