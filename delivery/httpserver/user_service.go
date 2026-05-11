package httpserver

import (
	"net/http"

	"github.com/sahar-mirtalebi/quiz-battle/dto"
	"github.com/sahar-mirtalebi/quiz-battle/pkg/httpmessage"

	"github.com/labstack/echo/v4"
)

func (s Server) userLogin(c echo.Context) error {
	var req dto.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) userRegister(c echo.Context) error {
	var req dto.RegisterRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	resp, err := s.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

func (s Server) userProfile(c echo.Context) error {
	bearerToken := c.Request().Header["Authorization"]

	claims, err := s.authSvc.ParseJWT(bearerToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := s.userSvc.Profile(dto.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmessage.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)
}
