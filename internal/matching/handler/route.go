package matchinghandler

import (
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (h Handler) SetupRoutes(e *echo.Echo) {
	userGroup := e.Group("/matching")
	userGroup.POST("/add-to-waiting-list", h.addToWaitingList, echojwt.JWT(h.signKey))

}
