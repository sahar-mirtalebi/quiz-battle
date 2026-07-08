package backofficeuserhandler

import (
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/sahar-mirtalebi/quiz-battle/delivery/httpserver/middleware"
	"github.com/sahar-mirtalebi/quiz-battle/entity"
)

func (h Handler) SetUpBackofficeUserRoute(e *echo.Echo) {
	backofficeGroup := e.Group("/backoffice/users")
	backofficeGroup.GET("", h.listAllUsers, echojwt.JWT(h.signKey), middleware.Authorization(h.authzSvc, entity.UserListPermission))
}
